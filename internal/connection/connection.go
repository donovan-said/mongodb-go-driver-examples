/*
- https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Client
- https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect
- https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Client.Disconnect
- https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Client.Ping
*/

package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func CloseConn(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {
		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

// This is a user defined method that returns mongo.Client, context.Context,
// context.CancelFunc and error. mongo.Client will be used for further database
// operation. context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and resource associated
// with it.

func OpenConn(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here deadline will of
	// 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// mongo.Connect return mongo.Client method
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel, err
}

// This is a user defined method that accepts mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.

func PingConn(client *mongo.Client, ctx context.Context) error {

	err := client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(">> Pinged MongoDB Successfully!")
	}
	return nil
}
