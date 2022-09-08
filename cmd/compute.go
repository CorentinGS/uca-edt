package cmd

import (
	"context"
	"fmt"
	"github.com/corentings/uca-edt/pkg/core"
	"github.com/corentings/uca-edt/pkg/database"
	"github.com/corentings/uca-edt/pkg/utils"
	"github.com/spf13/cobra"
	"log"
)

var computeGFile string
var computeEFile string

var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute the EDT",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Load var from .env file
		utils.LoadVar()

		err := database.Connect(database.MongoURL)
		if err != nil {
			log.Panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Printf("Compute the edt with the group file %s and the edt file %s", computeGFile, computeEFile)
		core.ComputeStudentEDT(computeGFile, computeEFile)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer func() {
			fmt.Println("Disconnect")
			err := database.Mg.Client.Disconnect(context.TODO())
			if err != nil {
				return
			}
		}()
	},
}
