package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CloseConn(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

func OpenConn(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here deadline will of
	// 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected successfully!")
	}

	return client, ctx, cancel, err
}

func PingConn(client *mongo.Client, ctx context.Context) error {

	err := client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Pinged to MongoDB Successfully!")
	}
	return nil
}
