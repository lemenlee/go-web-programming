package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "hello world",
		Author: Author{
			Id:   2,
			Name: "lemon",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "Have a greate day",
				Author:  "Adam",
			},
			Comment{
				Id:      4,
				Content: "Nice",
				Author:  "soojin",
			},
		},
	}

	// output, err := json.MarshalIndent(&post, "", "\t\t")
	// if err != nil {
	// 	fmt.Println("error")
	// 	return
	// }
	// err = ioutil.WriteFile("post.json", output, 0644)
	// if err != nil {
	// 	fmt.Println("Erro writting json")
	// 	return
	// }

	//使用encoder
	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("error create")
		return
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		return
	}
}
