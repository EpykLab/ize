/*
Copyright © 2024 EpykLab contact@epyklab.com
*/
package cmd

import (
	"github.com/EpykLab/ize/cmd/entries"
	"github.com/spf13/cobra"
)

// searchs3Cmd represents the searchs3 command
var searchs3Cmd = &cobra.Command{
	Use:   "searchs3",
	Short: "s3 objects in s3",
	Long: `searches objects that have been uploaded to the 
	s3 bucket defined in config file.
	
	Defaults to listing all objects in the defined bucket`,
	Run: func(cmd *cobra.Command, args []string) {
		filter := cmd.Flag("filter").Value.String()

		switch {
		case filter != "":
			entries.ListBucketObjects(filter)
		default:
			entries.ListBucketObjects("all")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchs3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchs3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchs3Cmd.Flags().String("filter", "", "filter for object in bucket")
}
