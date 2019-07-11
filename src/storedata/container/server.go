package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello", Author: "lemon"}
	post2 := Post{Id: 2, Content: "World", Author: "Sau sheong"}
	post3 := Post{Id: 3, Content: "lemon", Author: "Sau sheong"}
	post4 := Post{Id: 4, Content: "kv", Author: "lemon"}
	post5 := Post{Id: 5, Content: "ekasbo", Author: "abc"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)
	store(post5)

	fmt.Println(PostById[2])
	fmt.Println(PostById[1])

	for _, post := range PostByAuthor["Sau sheong"] {
		fmt.Println(post.Author)
	}

	for _, post := range PostByAuthor["lemon"] {
		fmt.Println(post.Content)
	}
}
