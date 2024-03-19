package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuanrenc/letCodeSpeakForItself/config"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "simpleDemo",
	Short: "a simple demo",
	Long:  "a simple demo to show my skills",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	cfg = &config.Config{}
	cfg.DbUser = conf.DbUser
	cfg.DbName = conf.DbName
	cfg.DbHost = conf.DbHost
	cfg.DbPassword = conf.DbPassword
	cfg.DbHost = conf.DbHost
	cfg.DbPort = conf.DbPort
	cfg.ListenSpec = conf.ListenHost + ":" + conf.ListenPort
}
