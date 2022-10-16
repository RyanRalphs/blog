package structs

//create struct for posts
type Post struct {
	ID      int32    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}


