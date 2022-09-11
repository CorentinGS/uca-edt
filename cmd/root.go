package cmd

import (
	"context"
	"fmt"
	"github.com/corentings/uca-edt/app/routes"
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/corentings/uca-edt/pkg/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "uca-edt",
	Version: "0.1.2",
	Short:   "uca-edt is a CLI to manage the UCA EDT",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Load var from .env file
		utils.LoadVar()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Connect(database.MongoURL)
		if err != nil {
			log.Panic(err)
		}

		defer func() {
			fmt.Println("Disconnected from the database")
			err := database.Mg.Client.Disconnect(context.TODO())
			if err != nil {
				return
			}
		}()
		// Create the app
		app := routes.New()
		log.Panic(app.Listen(":3000"))
	},
}

func initFlags() {
	computeCmd.Flags().StringVarP(&computeGFile, "gfile", "g", "", "The group file to use")
	computeCmd.Flags().StringVarP(&computeEFile, "efile", "e", "", "The edt file to use")
	err := computeCmd.MarkFlagRequired("gfile")
	if err != nil {
		return
	}
	err = computeCmd.MarkFlagRequired("efile")
	if err != nil {
		return
	}
}

// Execute executes the root command. It is called by main.main().
func Execute() {
	initFlags()                    // Initialize the flags
	rootCmd.AddCommand(computeCmd) // Add the compute command
	if err := rootCmd.Execute(); err != nil {
		log.Println(err) // Print the error
		os.Exit(1)       // Exit with error
	}
}
