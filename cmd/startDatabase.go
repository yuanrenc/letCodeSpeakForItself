package cmd

import (
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
		database.CreateTable(db)
		database.InsertData(db)
		database.GetTasks(db)

	},
}

func init() {
	rootCmd.AddCommand(dataCmd)
}

// func getDataBaseConfig{}()
