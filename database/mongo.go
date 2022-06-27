package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb+srv://db_practices_root:db_practices_root@cluster0.qpg6q7s.mongodb.net/?retryWrites=true&w=majority"

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
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

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.
func Connect() (*mongo.Client, context.Context, context.CancelFunc, error, *mongo.Collection) {
	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	// Crear la referencia a las bases de datos en mongoDb.
	authentication_go := client.Database("authentication-go")
	auth_collection := authentication_go.Collection("auth")

	return client, ctx, cancel, err, auth_collection
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {
	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("connected successfully")
	return nil
}

func AddUser(userCollection *mongo.Collection, ctx context.Context, user interface{}) (*mongo.InsertOneResult, error) {
	result, err := userCollection.InsertOne(ctx, user)
	return result, err
}
