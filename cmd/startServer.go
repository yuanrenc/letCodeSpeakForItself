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
	listenSpec := cfg.ListenSpec
	listener, err := net.Listen("tcp", listenSpec)
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
	fmt.Printf("Server started at %v\n", listenSpec)
	http.HandleFunc("/data", dataHandler)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}

	go server.Serve(listener)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// get all tasks from database
	dbConfig := database.DatabaseConfig{
		DbUser:     cfg.DbUser,
		DBPassword: cfg.DbPassword,
		DbPort:     cfg.DbPort,
		DbName:     cfg.DbName,
		DbHost:     cfg.DbHost,
	}
	db := database.ConnectToDataBase(&dbConfig)
	defer db.Close()
	data := database.GetTasks(db)
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
