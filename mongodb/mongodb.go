package mongodb

import (
	"context"
	"entities"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var URL string = os.Getenv("MONGO_URL")

var collectionName string = os.Getenv("COLLECTION_NAME")

var database string = os.Getenv("DB_MONGO")

var ctx context.Context = context.Background()

func getConnexion() (*mongo.Database, error) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))

	if err != nil {
		return nil, err
	}

	return client.Database(database), nil
}

func GetStudents() (*[]entities.Students, error) {

	var res []entities.Students

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
