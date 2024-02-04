/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ize",
	Short: "Red team information collaboration framework (pronounced like 'eyes')",
	Long: `
	ize (pronounced like "eyes") is a collaborative red teaming tool 
	designed to streamline artifact collection and facilitate real-time 
	information sharing among red team members. With ize, red teams can 
	securely store AWS STS credentials and other critical findings in a 
	shared database, enabling seamless collaboration. The tool also 
	simplifies the upload of screenshots, videos, logs, and other artifacts 
	to object storage, ensuring centralized access for reporting and 
	future reference. ize prioritizes security and efficiency, offering 
	robust encryption and access controls while empowering red teams to 
	focus on their objectives. Elevate your offensive security operations 
	with ize and enhance your team's effectiveness today.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ize.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ize" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ize")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		fmt.Println("\n")
	}
}
