/*
Copyright © 2024 EpykLab contact@epyklab.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"slices"

	entries "github.com/EpykLab/ize/cmd/entries"
	tags "github.com/EpykLab/ize/cmd/tags"
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
		tag := cmd.Flag("tag").Value.String()
		name := cmd.Flag("name").Value.String()
		author := cmd.Flag("author").Value.String()

		switch {

		case tag != "":
			present := slices.Contains(tags.Tag, tag)

			if present != true {
				log.Fatal("supplied tag not currently supported")
				os.Exit(1)
			} else {
				entries.SearchEntriesByTag(tag)
			}
		case name != "":
			entries.SearchEntriesByName(name)

		case author != "":
			entries.SearchEntriesByAuthor(author)

		default:
			fmt.Println("no search tags provided")
		}

	},
}

func init() {
	rootCmd.AddCommand(searchdbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//searchdbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchdbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchdbCmd.Flags().String("tag", "", "tag to search entries by")
	searchdbCmd.Flags().String("name", "", "name of the entry to search by")
	searchdbCmd.Flags().String("author", "", "author to search entries by")
}
