/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchs3Cmd represents the searchs3 command
var searchs3Cmd = &cobra.Command{
	Use:   "searchs3",
	Short: "s3 objects in s3",
	Long: `searches objects that have been uploaded to the 
	s3 bucket defined in config file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("searchs3 called")
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
	// searchs3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
