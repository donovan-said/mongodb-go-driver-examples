/*
This code is primarily derived from:
- https://www.mongodb.com/docs/drivers/go/current/usage-examples/
- https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
- https://www.geeksforgeeks.org/how-to-use-go-with-mongodb/
*/

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/donovan-said/mongodb-go-driver-examples/internal/connection"
	"github.com/donovan-said/mongodb-go-driver-examples/internal/find"
	"github.com/donovan-said/mongodb-go-driver-examples/internal/insertion"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MongoDB_URI = "mongodb://root:rootpassword@127.0.0.1:27017/"

// Struct is used for sample dataset
type Film struct {
	Name     string
	Year     int
	Genre    string
	Language string
}

func insertOneSample(client *mongo.Client, ctx context.Context) {
	// TODO Review the followuing:
	// Create  a object of type interface to  store the bson values, that
	// we are inserting into database.
	// var document interface{}
	// document := bson.D{{"Dune", 2022}, {"TMNT", 2023}}

	entry := Film{Name: "Dune", Year: 2020, Genre: "Science Fiction", Language: "English"}

	// InsertOne() accepts client , context, database name collection name and
	// an interface that will be inserted into the  collection. insertOne
	// returns an error and a result of insert in a single document into the
	// collection.
	insertOneResult, err := insertion.InsertOne(
		client, ctx, "entertainment", "films", entry,
	)

	// Handle the error
	if err != nil {
		panic(err)
	}

	// Print the insertion id of the document, if it is inserted.
	fmt.Println(">> Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)
}

func insertManySample(client *mongo.Client, ctx context.Context) {
	entries := []interface{}{
		Film{Name: "Deadpool", Year: 2024, Genre: "Super Hero", Language: "English"},
		Film{Name: "TMNT", Year: 2023, Genre: "Super Hero", Language: "English"},
		Film{Name: "Star Wars", Year: 2020, Genre: "Science Fiction", Language: "English"},
	}

	insertManyResult, err := insertion.InsertMany(
		client, ctx, "entertainment", "films", entries,
	)

	// Handle the error
	if err != nil {
		panic(err)
	}

	// Print the insertion ids of the documents, if they is inserted.
	fmt.Println(">> Result of InsertMany")

	for _, id := range insertManyResult.InsertedIDs {
		fmt.Println(id)
	}
}

func findManySample(client *mongo.Client, ctx context.Context) {
	// create a filter an option of type interface, that stores bjson
	// objects.
	// var filter, option interface{}

	// filter := Film{Name: "TMNT", Year: 2023, Genre: "Super Hero", Language: "English"}

	filter := bson.D{{"name", "TMNT"}}

	cursor, err := find.Query(client, ctx, "entertainment", "films", filter)

	if err != nil {
		panic(err)
	}

	var results []Film
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}

func userPrompt() (response string) {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println(">> Do you want to populate the database? (yes/no): ")
	name, _ := scanner.ReadString('\n')

	return strings.TrimSpace(name)

}

func main() {

	//----------------------------------------------------------------------
	// User input
	populate := userPrompt()

	//----------------------------------------------------------------------
	// Establish connection

	// Get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := connection.OpenConn(MongoDB_URI)
	if err != nil {
		panic(err)
	}

	// Release resource when the main function is returned.
	defer connection.CloseConn(client, ctx, cancel)

	// Ping mongoDB with Ping method
	connection.PingConn(client, ctx)

	//----------------------------------------------------------------------
	// Insertion switch statement

	switch populate {
	case "yes":
		fmt.Printf(">> Populating the MongoDB database!\n")
		// Insert One
		insertOneSample(client, ctx)

		// Insert Many
		insertManySample(client, ctx)
	case "no":
		fmt.Printf(">> Not populating the MongoDB database!\n")
	}

	//----------------------------------------------------------------------
	// Find

	findManySample(client, ctx)

}
