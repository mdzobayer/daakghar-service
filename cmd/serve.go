package cmd

import (
	"fmt"
	"net/http"

	"github.com/daakghar-service/config"
	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/routes"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.Get().Load()
		conn.DB().Connect(config.Get().Mgo())

		if conn.DB().Err() != nil {
			return conn.DB().Err()
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server starting at port", config.Get().Basic().Port)

		http.ListenAndServe(fmt.Sprintf(":%s", config.Get().Basic().Port), routes.Get())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
