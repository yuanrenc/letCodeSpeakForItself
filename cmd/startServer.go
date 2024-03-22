package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	listenSpec := cfg.ListenHost + ":" + cfg.ListenPort
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
	go server.Serve(listener)
	waitForSignal()
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// get all tasks from database
	enableCors(&w)
	dbConfig := database.DatabaseConfig{
		DbUser:     &cfg.DbUser,
		DbPassword: &cfg.DbPassword,
		DbPort:     &cfg.DbPort,
		DbName:     &cfg.DbName,
		DbHost:     &cfg.DbHost,
	}
	db := database.ConnectToDatabase(&dbConfig)
	defer db.Close()
	data := database.GetTasks(db)
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func waitForSignal() {
	xsig := make(chan os.Signal)
	signal.Notify(xsig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	hsig := make(chan os.Signal)
	signal.Notify(hsig, syscall.SIGHUP)
	for {
		select {
		case s := <-xsig:
			log.Fatalf("Got signal: %v, exiting.", s)
		case s := <-hsig:
			log.Printf("Got signal: %v, continue.", s)
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
