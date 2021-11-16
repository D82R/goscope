package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/d82r/goscope/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var allProgs bool

// pipeCmd represents the pipe command
var pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Print only wildcard domains to stdout for program",
	Long:  `
GoScope v1.0
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// set default database name if not specified through cli
		if database == "" {
			database = viper.GetString("config.database")
		}

		// establish database connection
		db := data.ConnectDb(database)
		defer db.Close()

		if len(progName) < 2 && !allProgs {
			fmt.Println("Must specify a program name")
			os.Exit(1)
		}

		progName = strings.ToLower(progName)

		data.PipeCommand(db, progName, allProgs)
	},
}

func init() {
	
	rootCmd.AddCommand(pipeCmd)
	pipeCmd.Flags().BoolVarP(&allProgs, "all", "a", false, "Print all wildcard domains for all programs to stdout")
	pipeCmd.Flags().StringVarP(&progName, "program", "p", "", "program name")
}
