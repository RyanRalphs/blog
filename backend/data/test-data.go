package data

import (
	"fmt"

	"github.com/ryanralphs/blog/structs"
)

func FakePosts(count int) []structs.Post {
	posts := make([]structs.Post, count)
	for i := 0; i < count; i++ {
		posts[i] = structs.Post{
			ID:      i,
			Title:   fmt.Sprintf("Post %d", i),
			Content: fmt.Sprintf("Post %d content", i),
		}
	}
	return posts
}
