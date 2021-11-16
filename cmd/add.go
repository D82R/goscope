package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/d82r/goscope/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var burpfile string
var database string
var progName string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new program",
	Long: `
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

		if len(progName) < 2 {
			fmt.Println("Must specify a program name")
			os.Exit(1)
		}

		progName = strings.ToLower(progName)

		// parse burp file
		inscope, outscope := data.ParseJson(burpfile)
		// clean parsed data
		inscope, outscope = data.CleanDomains(inscope, outscope)

		// add program to database
		data.AddProgram(db, progName, inscope, outscope)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&burpfile, "burpfile", "b", "", "specify Burp scope file")
	addCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "set database name (default: ./scope.db)")
	addCmd.Flags().StringVarP(&progName, "program", "p", "", "program name")
}
