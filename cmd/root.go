package cmd

import (
	"fmt"
	"os"

	"github.com/gabey14/tasker-gocli/taskercli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:     "tasker cli",
		Short:   "tasker cli using Go, simple task management for command line",
		Version: "0.0.1"}
)

// Execute - execute is the main entry point for the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// non 0 code indicates an error
		os.Exit(1)
	}
}

func initConfig() {
	var configFile string
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home := taskercli.UserHomeDir()

		viper.AddConfigPath(home)
		viper.AddConfigPath(home + "/.config/taskercli")
		viper.SetConfigName(".taskercli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error: Can't load config file:", viper.ConfigFileUsed())
		fmt.Println("Run 'taskercli --help' for usage")
		os.Exit(1)
	}
}
