/*
This code is primarily derived from https://www.geeksforgeeks.org/how-to-use-go-with-mongodb/
*/

package main

import (
	"context"
	"fmt"

	"github.com/donovan-said/mongodb-go-driver-examples/internal/connection"
	"github.com/donovan-said/mongodb-go-driver-examples/internal/find"
	"github.com/donovan-said/mongodb-go-driver-examples/internal/insertion"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MongoDB_URI = "mongodb://root:rootpassword@127.0.0.1:27017/"

type Film struct {
	Name string
	Year int
}

func insertOneSample(client *mongo.Client, ctx context.Context) {
	// TODO Review the followuing:
	// Create  a object of type interface to  store the bson values, that
	// we are inserting into database.
	// var document interface{}
	// document := bson.D{{"Dune", 2022}, {"TMNT", 2023}}

	entry := Film{Name: "Dune", Year: 2022}

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
	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)
}

func insertManySample(client *mongo.Client, ctx context.Context) {
	entries := []interface{}{
		Film{Name: "Deadpool", Year: 2024},
		Film{Name: "TMNT", Year: 2023},
	}

	insertManyResult, err := insertion.InsertMany(
		client, ctx, "entertainment", "films", entries,
	)

	// Handle the error
	if err != nil {
		panic(err)
	}

	// Print the insertion ids of the documents, if they is inserted.
	fmt.Println("Result of InsertMany")

	for _, id := range insertManyResult.InsertedIDs {
		fmt.Println(id)
	}
}

func findSample(client *mongo.Client, ctx context.Context) {
	// create a filter an option of type interface, that stores bjson
	// objects.
	// var filter, option interface{}

	// filter  gets all document, with maths field greater that 70
	filter := Film{Name: "Dune", Year: 2022}

	cursor, err := find.Query(client, ctx, "entertainment", "films", filter)

	if err != nil {
		panic(err)
	}

	var results []Film
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}

func main() {

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
	// Insert One

	insertOneSample(client, ctx)

	//----------------------------------------------------------------------
	// Insert Many

	insertManySample(client, ctx)

	//----------------------------------------------------------------------
	// Find

	findSample(client, ctx)

}
