package insertion

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// InsertOne is a user defined method, used to insert documents into collection
// returns result of InsertOne and error if any.
func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	// select database and collection ith Client.Database method and
	// Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

// InsertMany is a user defined method, used to insert documents into collection
// returns result of InsertOne and error if any.
func InsertMany(client *mongo.Client, ctx context.Context, dataBase, col string, doc []interface{}) (*mongo.InsertManyResult, error) {
	// select database and collection ith Client.Database method and
	// Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertMany accept two argument of type Context and of empty interface
	result, err := collection.InsertMany(ctx, doc)
	return result, err
}
