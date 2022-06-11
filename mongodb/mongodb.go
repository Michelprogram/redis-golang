package mongodb

import (
	"context"
	"entities"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//- MONGO_URL=mongodb://api_user:api1234@localhost:27017/?authMechanism=DEFAULT&authSource=api_prod_db
//mongodb://api_user:api1234@localhost:27017/?authMechanism=DEFAULT&authSource=api_prod_db
var URL string = os.Getenv("MONGO_URL")

//var URL = "mongodb://api_user:api1234@localhost:27017/?authMechanism=SCRAM-SHA-1&authSource=api_prod_db"
var ctx context.Context = context.Background()

//var database string = "Redis"

var database string = "api_prod_db"

func getConnexion() (*mongo.Database, error) {

	fmt.Println("Conexion :", URL)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))

	if err != nil {
		return nil, err
	}

	return client.Database(database), nil
}

func GetStudents() (*[]entities.Students, error) {

	var res []entities.Students

	var collectionName string = "students"

	conn, err := getConnexion()

	if err != nil {
		return nil, err
	}

	defer conn.Client().Disconnect(ctx)

	var collection *mongo.Collection = conn.Collection(collectionName)

	contactCursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	if err = contactCursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
