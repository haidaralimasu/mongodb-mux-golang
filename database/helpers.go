package database

import (
	"context"

	"github.com/Haidar1528/mongodb_mux_golang/helper"
	"github.com/Haidar1528/mongodb_mux_golang/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// insert one record
func InsertOneMovie(movie model.Netflix) interface{} {
	inserted, err := Collection.InsertOne(context.Background(), movie)
	helper.HandleErr(err)

	return inserted.InsertedID
}

// update one record
func UpdateOneMovie(movieId string) int64 {
	id, err := primitive.ObjectIDFromHex(movieId)
	helper.HandleErr(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := Collection.UpdateOne(context.Background(), filter, update)
	helper.HandleErr(err)

	return result.ModifiedCount
}

// delete one record
func DeleteOneMovie(movieId string) int64 {
	id, err := primitive.ObjectIDFromHex(movieId)
	helper.HandleErr(err)

	filter := bson.M{"_id": id}
	deleteCount, err := Collection.DeleteOne(context.Background(), filter, nil)
	helper.HandleErr(err)

	return deleteCount.DeletedCount
}

// delete all records
func DeleteAllMovies() int64 {
	deleteResult, err := Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	helper.HandleErr(err)

	return deleteResult.DeletedCount
}

func GetAllMovies() []primitive.M {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	helper.HandleErr(err)

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		helper.HandleErr(err)

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}
