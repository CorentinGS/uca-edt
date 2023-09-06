package cmd

import (
	"context"
	"log"

	"github.com/corentings/uca-edt/pkg/core"
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/corentings/uca-edt/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	computeGFile string // The group file to use
	computeEFile string // The edt file to use
)

// computeCmd represents the compute command
var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute the EDT",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Load var from .env file
		utils.LoadVar()

		// Connect to the database
		err := database.Connect(database.MongoURL)
		if err != nil {
			log.Panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Compute the edt
		log.Printf("Compute the edt with the group file %s and the edt file %s", computeGFile, computeEFile)

		// Compute the edt
		core.ComputeStudentEDT(computeGFile, computeEFile)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// Disconnect from the database
		defer func() {
			log.Println("Disconnect")
			err := database.Mg.Client.Disconnect(context.TODO())
			if err != nil {
				log.Panic(err)
				return
			}
		}()
	},
}
