/*
Copyright © 2024 EpykLab contact@epyklab.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"slices"

	tags "github.com/EpykLab/ize/cmd/tags"
	userio "github.com/EpykLab/ize/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "create and entry in MongoDB",
	Long: `Creates an entry in the mongo database defined
	in the config file.
	
	NOTE: Search functionality relies heavily on tagging entries. While
	tagging is not enforced, it is highly encouraged to ensure
	documents are easier to find by you and others`,
	Run: func(cmd *cobra.Command, args []string) {
		val := cmd.Flag("tag").Value.String()
		author := viper.GetString("team.self")
		present := slices.Contains(tags.Tag, val)
		if val == "" {
			fmt.Println("No tag name supplied")
			os.Exit(1)
		}
		if present != true {
			log.Fatal("supplied tag not currently support")
			os.Exit(1)
		} else {
			userio.SetEntryTypeByTag(val, author)
		}
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	entryCmd.Flags().String("tag", "", "entry type. automatically applies tag based on type")

}
