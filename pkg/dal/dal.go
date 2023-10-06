package dal

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MustGetClient gets a mongo client.
// If the client cannot be created, the program will exit.
func MustGetClient(ctx context.Context) *mongo.Client {
	var (
		err    error
		client *mongo.Client
	)

	connectionString := viper.Get("mongo.uri").(string)

	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to the database
	if client, err = mongo.Connect(ctx, clientOptions); err != nil {
		log.Fatal(err)
	}

	// Note: defer is not needed here because the client is used throughout the program.
	//defer func(client *mongo.Client, ctx context.Context) {
	//	err := client.Disconnect(ctx)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(client, ctx)

	// Ping the database
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return client
}
