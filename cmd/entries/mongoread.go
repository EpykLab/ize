package entries

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchEntriesByTag(tag string) error {
	// Define MongoDB connection URI
	connstr := viper.GetString("mongo_conn_str")
	collection := viper.GetString("collection")
	database := viper.GetString("database")

	// Set client options
	clientOptions := options.Client().ApplyURI(connstr)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	// Access the database and collection
	col := client.Database(database).Collection(collection)

	filter := bson.M{"tags": tag}

	cursor, err := col.Find(context.Background(), filter)
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	// Iterate over the cursor and print document contents
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("Error decoding document:", err)
		}

		// Convert the map to a JSON string with indentation
		jsonData, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatal("Error marshaling JSON:", err)
		}

		// Print the pretty printed JSON string
		fmt.Println(string(jsonData))
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("Error iterating over cursor:", err)
	}

	return nil

}
