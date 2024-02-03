/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ize",
	Short: "Red team information collaboration framework (pronounced like 'eyes')",
	Long: `ize - Collaborative Red Teaming Tool

	ize (pronounced like "eyes") is a powerful command-line tool designed 
	for red teams to enhance collaboration and streamline artifact collection 
	during engagements. With ize, red team members can efficiently share critical 
	information and securely store artifacts for centralized access and reporting.
	
	Key Features:
	- Collaborative Information Sharing: ize enables red team members 
	to share crucial findings and intelligence seamlessly. For instance, 
	if one team member discovers AWS STS credentials or other important 
	information, they can easily save it to the shared database using ize. 
	This facilitates real-time collaboration and ensures that valuable 
	insights are accessible to the entire team.
	
	- Centralized Artifact Collection: ize simplifies the process of 
	collecting and managing artifacts such as screenshots, videos, logs, 
	and other relevant data. Red team members can upload these artifacts 
	directly to object storage via ize, ensuring that all essential 
	information is securely stored and organized for future reference 
	and reporting.
	
	- Secure and Efficient: With ize, security and efficiency are paramount. 
	The tool prioritizes data security by implementing robust encryption 
	and access control mechanisms. Additionally, ize's intuitive interface 
	and streamlined workflows empower red teams to focus on their objectives 
	without being bogged down by cumbersome processes.
	
	Whether you're conducting penetration tests, red team exercises, or 
	security assessments, ize is your trusted companion for collaborative 
	information sharing and artifact management. Empower your red team with 
	ize and elevate your offensive security operations to new heights.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ize.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
