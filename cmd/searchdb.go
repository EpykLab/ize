/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"slices"

	entries "github.com/Epyklab/ize/cmd/entries"
	tags "github.com/Epyklab/ize/cmd/tags"
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
		val := cmd.Flag("by").Value.String()
		present := slices.Contains(tags.Tag, val)
		if val == "" {
			fmt.Println("No tag name supplied")
			os.Exit(1)
		}
		if present != true {
			log.Fatal("supplied tag not currently supported")
			os.Exit(1)
		} else {
			entries.SearchEntriesByTag(val)
		}
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
	searchdbCmd.Flags().String("by", "", "tag to search entries by")
}
