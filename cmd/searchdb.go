/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchdbCmd represents the searchdb command
var searchdbCmd = &cobra.Command{
	Use:   "searchdb",
	Short: "searches documents stored in MongoDB",
	Long: `search for document entries written to mongodb.
	
	NOTE: Search functionality relies heavily on tagging entries. While
	tagging is not enforced, it is highly encouraged to ensure
	documents are easier to find by you and others`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("searchdb called")
	},
}

func init() {
	rootCmd.AddCommand(searchdbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchdbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchdbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
