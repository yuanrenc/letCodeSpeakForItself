package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/yuanrenc/letCodeSpeakForItself/database"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("starting server...\n")
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func startServer() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)
	fmt.Printf("Server started at http://localhost:8080\n")
	http.HandleFunc("/data", dataHandler)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectToDataBase()
	defer db.Close()
	data := database.GetTasks(db)
	fmt.Print(data[0].Title)
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
