/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
package entries

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertJSONData inserts JSON data into MongoDB
func MakeMongoEntry(jsonData []byte) error {
	// Define MongoDB connection URI
	connstr := viper.GetString("database.mongo_conn_str")
	collection := viper.GetString("database.collection")
	database := viper.GetString("database.database")

	// Set client options
	clientOptions := options.Client().ApplyURI(connstr)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Disconnect from MongoDB after the function completes
	defer client.Disconnect(context.Background())

	// Access the database and collection
	col := client.Database(database).Collection(collection)

	// Parse JSON data
	var data interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}

	// Insert the JSON data into MongoDB
	_, err = col.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}

	fmt.Println("Data inserted successfully into MongoDB")

	return nil
}
