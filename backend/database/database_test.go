package database

import (
	"testing"

	"github.com/ryanralphs/blog/structs"
)

var TestPosts = []structs.Post{
	{
		ID: 1,
		Title: "Test Post",
		Content: "Test Content",
	},
	{
		ID: 2,
		Title: "Test Post Two",
		Content: "Test Content Two",
	},

	}

func TestConnectDB(t *testing.T) {
		collection, ctx, err := ConnectDB("post-tests")
		if err != nil {
			t.Errorf("Error connecting to database: %v", err)
		}
		if collection == nil {
			t.Errorf("Collection does not exist! Connect to your local MongoDB instance and try again.")
		}
		if ctx == nil {
			t.Errorf("May be having context issues")
		}
}

func TestStorePost(t *testing.T){
	TestPost := TestPosts[0]
	collection, ctx,err := ConnectDB("post-tests")
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}

	_, err = StorePost(TestPost, collection, ctx)
	if err != nil {
		t.Errorf("Error storing post: %v", err)
	}
}


func TestGetPostById(t *testing.T){
	TestPost := TestPosts[0]
	collection, ctx, err := ConnectDB("post-tests") 
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	_, err = StorePost(TestPost, collection, ctx)
	if err != nil {
		t.Errorf("Error storing post: %v", err)
	}
	res, err := GetPostById(TestPost.ID, collection, ctx)

	if err != nil {
		t.Errorf("Error storing post: %v", err)
	}
	if(res["id"] != TestPost.ID){
		t.Errorf("Post ID does not match")
	}
}

func TestGetAllPosts(t *testing.T){

	collection, ctx, err := ConnectDB("post-tests")
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	for _, post := range TestPosts {
	_, err := StorePost(post, collection, ctx)
	if err != nil {
		t.Errorf("Error storing post: %v", err)
	}
}
	res, err := GetAllPosts(collection, ctx)

	if err != nil {
		t.Errorf("Error storing post: %v", err)
	}
	if(len(res) < len(TestPosts)){
		t.Errorf("Not all posts were returned")
	}
}


func TestUpdatePost(t *testing.T){
	collection, ctx,err := ConnectDB("post-tests")
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}

	TestPost := structs.Post{ID: 1, Title: "Test Post update", Content: "Test Content update"}
	
	res, err := UpdatePost(int(TestPost.ID), TestPost, collection, ctx)

	if err != nil {
		t.Errorf("Error storing post: %v", err)
	}

	if res.MatchedCount != 1 {
		t.Errorf("Post was not updated")
	}
}

func TestDeletePost(t *testing.T){
	collection, ctx,err := ConnectDB("post-tests")
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}

	TestPost := structs.Post{ID: 1, Title: "Test Post update", Content: "Test Content update"}
	
	res, err := DeletePost(int(TestPost.ID), collection, ctx)

	if err != nil {
		t.Errorf("Error deleting post: %v", err)
	}

	if res.DeletedCount != 1 {
		t.Errorf("Post was not deleted")
	}

}