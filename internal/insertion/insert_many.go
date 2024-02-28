// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany

package insertion

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

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
