package cmd

import (
	"log"
	"os"

	"github.com/rhasan33/goplate/config"
	"github.com/rhasan33/goplate/conn"
	"github.com/spf13/cobra"
)

// RootCmd declares the command information
var RootCmd = &cobra.Command{
	Use:   "reader",
	Short: "Book reading platform",
	Long:  "Book reading platform",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	config.Init()
	conn.ConnectDB()
	conn.ConnectRedis()
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
