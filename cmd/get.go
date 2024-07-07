/*
Copyright © 2024 EpykLab contact@epyklab.com
*/

package cmd

import (
	"github.com/EpykLab/ize/cmd/entries"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "gets3object",
	Short: "Download object from S3 storage",
	Long: `Download an object from S3 compatable storage.

	Takes a filename as an argument and will seek to download from the s3 storage specified in S3`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := cmd.Flag("object").Value.String()
		entries.DownloadFromS3(filepath)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().String("object", "", "Download object from s3")
}
