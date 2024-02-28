// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.Find

package find

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// query is user defined method used to query MongoDB, that accepts
// mongo.client,context, database name, collection name, a query and field.

// database name and collection name is of type string. query is of type
// interface. field is of type interface, which limits the field being returned.

// query method returns a cursor and error.
func Find(client *mongo.Client, ctx context.Context, dataBase, col string, filter interface{}) (result *mongo.Cursor, err error) {

	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find, that returns a mongo.cursor based on
	// query and field.
	result, err = collection.Find(ctx, filter)

	return
}
