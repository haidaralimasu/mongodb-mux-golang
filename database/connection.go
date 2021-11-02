package database

import (
	"context"
	"fmt"

	"github.com/Haidar1528/mongodb_mux_golang/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString := helper.GoDotEnvVariable(CONNECTION_STRING)
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var Collection *mongo.Collection

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	helper.HandleErr(err)
	fmt.Println("Sucessfully connected with MongoDB")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connection is ready")
}
