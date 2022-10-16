package database

import (
	"context"
	"fmt"
	"os"

	"github.com/ryanralphs/blog/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
func ConnectDB(collectionName string) (*mongo.Collection, context.Context, error)  {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetDirect(true))
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	collection := client.Database("posts").Collection(collectionName)

	
	return collection, ctx, nil
}

func StorePost(request structs.Post, collection *mongo.Collection, ctx context.Context) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(ctx, request)
	if err != nil {
		message := fmt.Errorf("Error inserting post: %v", err)
		return nil, message
	}
	return res, nil
}

func GetPostById(id int32, collection *mongo.Collection, ctx context.Context) (primitive.M, error) {
	res := collection.FindOne(ctx, bson.M{"id": id})
	var result bson.M
	if err := res.Decode(&result); err != nil {
		message := fmt.Errorf("Error getting post: %v", err)
		return nil, message
	}
	return result, nil
}

func GetAllPosts(collection *mongo.Collection, ctx context.Context) ([]primitive.M, error) {
	res, err := collection.Find(ctx, bson.M{})
	if err != nil {
		message := fmt.Errorf("Error getting posts: %v", err)
		return nil, message
	}
	var results []bson.M
	if err = res.All(ctx, &results); err != nil {
		message := fmt.Errorf("Error getting posts: %v", err)
		return nil, message
	}
	return results, nil
}

func UpdatePost(id int, request structs.Post, collection *mongo.Collection, ctx context.Context) (*mongo.UpdateResult, error) {
	res, err := collection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": request})
	if err != nil {
		message := fmt.Errorf("Error updating post: %v", err)
		return nil, message
	}
	
	return res, nil
}

func DeletePost(id int, collection *mongo.Collection, ctx context.Context) (*mongo.DeleteResult, error) {
	res, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		message := fmt.Errorf("Error deleting post: %v", err)
		return nil, message
	}

	return res, nil
}


