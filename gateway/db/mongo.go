package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ddld93/tsrims/auth/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string

func init() {
	cfg := utils.LoadEvn()
	uri = cfg.MongoURI
}

func StartMongo() *mongo.Client {
    // Create a context with a 10-second timeout
    ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
    fmt.Printf("mongodb uri >>> %s \n", uri)

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    log.Println("Connecting to mongo ......")
    if err != nil {
        log.Fatal("Error connecting to mongoDB", err)
    }
	err = client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(ctx)
		log.Fatal("Error testing MongoDB connection", err)
	}
	log.Println("Connected to mongoDB")
    return client
}

