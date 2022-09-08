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
	Version: "0.1.0",
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

func Execute() {
	initFlags()
	rootCmd.AddCommand(computeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
