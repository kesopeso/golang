package main

import (
	"embed"
	"fmt"
	"log"
	"templating"
)

//go:embed posts/*
var postsDir embed.FS

func main() {
	posts, err := templating.ReadPostsFromFS(postsDir, "posts")
	if err != nil {
		log.Fatal("Something went really wrong", err)
	}
	fmt.Println("These are all the posts", posts)
}
