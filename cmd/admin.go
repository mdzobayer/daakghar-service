package cmd

import (
	"fmt"
	"log"

	"github.com/daakghar-service/config"
	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/db"
	"github.com/daakghar-service/dbq"
	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "create admin user. password=admin, user name=admin",
	Long:  `create admin user. password=admin, user name=admin`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("admin called", config.Get().Basic().Name)

		role := db.AdminRole()
		dbr := dbq.NewRole(conn.DB().DB())
		err := dbr.Put(&role)
		if err != nil {
			log.Println("could not crate admin role", err)
			return
		}

		usr, err := db.AdminUser(role.ID)
		if err != nil {
			log.Println("could not generate admin user", err)
			return
		}

		dbu := dbq.NewUser(conn.DB().DB())
		err = dbu.Put(&usr)
		if err != nil {
			log.Println("could not create admin user", err)
			return
		}

		fmt.Println("admin user created")

	},
}

func init() {
	createCmd.AddCommand(adminCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
