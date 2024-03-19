package cmd

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/yuanrenc/letCodeSpeakForItself/database"
)

var dataCmd = &cobra.Command{
	Use:   "database",
	Short: "active database service: create table and insert data",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.ConnectToDataBase()
		defer db.Close()
		// if the table already exist, will not create a new table
		database.CreateTable(db)
		// if the data doesn't exist, will insert data
		if len(database.GetTasks(db)) == 0 {
			fmt.Println("Data doesn't exist, inserting data")
			database.InsertData(db)
		}
	},
}

func init() {
	rootCmd.AddCommand(dataCmd)
}

// func getDataBaseConfig{}()
