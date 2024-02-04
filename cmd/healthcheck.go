/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// healthcheckCmd represents the healthcheck command
var healthcheckCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "check connections to MongoDB and object storage",
	Long: `Checks whether connections to MongoDB and object store 
	are configured properly. Configurations are defined in 
	$HOME/.ike.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

		MONGO_URI := viper.GetString("mongo_conn_str")

		clientOptions := options.Client().ApplyURI(MONGO_URI)
		client, err := mongo.Connect(context.Background(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.Background(), nil)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connection Successful")
		}
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// healthcheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// healthcheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
