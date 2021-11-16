package cmd

import (
	"github.com/d82r/goscope/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a program from the database",
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

		// remove program
		data.RemoveProgram(db, progName)
	},
}

func init() {

	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&progName, "program", "p", "", "program name")
}
