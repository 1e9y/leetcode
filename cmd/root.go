package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "leetcode",
	Short: "LeetCode Kit is a collection of helpers for solving LeetCode problems",
	// TODO: Complete long description
	Long: ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags. Those flags are global in the application.
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.newApp.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Read config file from project directory
		viper.AddConfigPath(".")
		// Search config in home directory with name ".leetcode" (without extension).
		viper.SetConfigType("yaml")
		viper.SetConfigName(".leetcode")
	}

	// Read in environment variables that match.
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
