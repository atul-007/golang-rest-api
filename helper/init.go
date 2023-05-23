package helper

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://atulranjan789:atulranjan@cluster0.ebjtzlx.mongodb.net/?retryWrites=true&w=majority"
const dbname = "netflix"
const colname = "watchlist"

// most important
var collection *mongo.Collection

func Init() {
	clientoption := options.Client().ApplyURI(connectionstring)

	//connect with mongodb
	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection sucess")

	collection = client.Database(dbname).Collection(colname)

	//collection instance
	fmt.Println("collection instance is ready ")
}
