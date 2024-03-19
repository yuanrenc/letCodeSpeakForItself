package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"embed"

	"github.com/spf13/cobra"
	"github.com/yuanrenc/letCodeSpeakForItself/database"
)

//go:embed asset/index.html
var staticFiles embed.FS

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
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	listener, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		log.Printf("Error creating listener - %v\n", err)
	}
	file, err := staticFiles.ReadFile("asset/index.html")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(file)
	})
	http.HandleFunc("/data", dataHandler)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server started at http://localhost:8080\n")
	go server.Serve(listener)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// get all tasks from database
	db := database.ConnectToDataBase()
	defer db.Close()
	data := database.GetTasks(db)
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
