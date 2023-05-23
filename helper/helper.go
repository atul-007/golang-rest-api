package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/atul-007/golang-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	Init()
}

//mongodb helper  -file

//insert 1 record

func Insertonemovie(movie models.Netflix) interface{} {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted one movie in db with id:", inserted.InsertedID)
	return inserted.InsertedID
}

//update 1 record

func Updateonerecord(movieid string) {
	id, err := primitive.ObjectIDFromHex(movieid)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count :", result.ModifiedCount)
}

//delete 1  record

func Deleteonerecord(movieid string) {

	id, _ := primitive.ObjectIDFromHex(movieid) //to convert string to objectid
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie got deleted with movie count:", deletecount)
}

//delete all records from mongodb

func Deleteallrecords() int64 {

	deleteresult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) //pass no value as filter to select all the values

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("no of movies deleted :", deleteresult.DeletedCount)
	return deleteresult.DeletedCount
}

func Deletemovies() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.D{{}}) //to get all the values in the cursor

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M //similar to  bson.m but better in this case

	for cursor.Next(context.Background()) { //to get all the values in the cursor
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}
func Getallmovies() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.D{{}}) //to get all the values in the cursor

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M //similar to  bson.m but better in this case

	for cursor.Next(context.Background()) { //to get all the values in the cursor
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}
