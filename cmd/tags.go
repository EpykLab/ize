/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	tags "github.com/Epyklab/ize/cmd/tags"
	"github.com/spf13/cobra"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "list avaiable tags",
	Long: `List the available tags that can be used in 
	writing documents to the database
	
	NOTE: custom tags are not supported at this time`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, x := range tags.Tag {
			fmt.Println(x)
		}
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
