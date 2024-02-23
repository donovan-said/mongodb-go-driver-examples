// https://www.geeksforgeeks.org/how-to-use-go-with-mongodb/

package main

import (
	//Import local packages
	"github.com/donovan-said/mongodb-go-driver-examples/internal/connection"
)

const MongoDB_URI = "mongodb://root:rootpassword@127.0.0.1:27017/"

func main() {
	client, ctx, cancel, err := connection.OpenConn(MongoDB_URI)
	if err != nil {
		panic(err)
	}

	defer connection.CloseConn(client, ctx, cancel)
	connection.PingConn(client, ctx)
}
