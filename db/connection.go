package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoConn is the object to connect MongoDB */
var MongoConn = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://yovadiaz:yova1234@clusterapi.4n5n3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConnectDB is the func to connect to MongoDB */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la MongoDB")
	return client
}

/* CheckingConnection is the func to Ping MongoDB */
func CheckingConnection() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
