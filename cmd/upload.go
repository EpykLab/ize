/*
Copyright © 2024 EpykLab contact@epyklab.com
*/
package cmd

import (
	"github.com/EpykLab/ize/cmd/entries"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload asset to s3",
	Long: `Uploads object(s) to the s3 bucket that is defined in the 
	config filed"`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := cmd.Flag("file").Value.String()
		entries.UploadToS3(filepath)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	uploadCmd.Flags().String("file", "", "path to object to upload")
}
