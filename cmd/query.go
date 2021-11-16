package cmd

import (
	"strings"

	"github.com/d82r/goscope/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var filter bool
var list bool

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query database for program scope",
	Long: `
GoScope v1.0
`,
	Run: func(cmd *cobra.Command, args []string) {

		progName = strings.ToLower(progName)

		if database == "" {
			database = viper.GetString("config.database")
		}

		// establish database connection
		db := data.ConnectDb(database)
		defer db.Close()

		// list all programs in database
		if list {
			data.ListPrograms(db)
		}

		// query the database for a program and get the scope
		if progName != "" {
			data.QueryProgram(db, progName)
		}
	},
}

func init() {

	rootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringVarP(&database, "database", "d", "", "database location (default: ./scope.db)")
	queryCmd.Flags().BoolVarP(&list, "list", "l", false, "list all programs")
	queryCmd.Flags().StringVarP(&progName, "program", "p", "", "program name")
}
