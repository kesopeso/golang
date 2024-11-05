package templating_test

import (
	"reflect"
	"templating"
	"testing"
	"testing/fstest"
)

func TestReadPostsFromFS(t *testing.T) {
	postsFs := fstest.MapFS{
		"blogpost1.txt": {Data: []byte(`Title: Test blog post 1
Description: This is some description
Tags: test, post
---
Body body body`)},
		"blogpost2.txt": {Data: []byte(`Title: Test blog post 2
Description: This is another description
Tags: another, samesame, different
---
Trying out something
new line body test`)},
	}

	posts, err := templating.ReadPostsFromFS(postsFs, ".")

	if err != nil {
		t.Error("error should not happen", err)
	}

	if len(posts) != 2 {
		t.Errorf("blog posts count mismatch, wanted 2, got %d", len(posts))
	}

	post := templating.Post{Title: "Test blog post 1", Description: "This is some description", Tags: []string{"test", "post"}, Body: "Body body body"}
	if !containsPost(posts, post) {
		t.Errorf("post not found, all posts %v, search post %v", posts, post)
	}

	post = templating.Post{Title: "Test blog post 2", Description: "This is another description", Tags: []string{"another", "samesame", "different"}, Body: `Trying out something
new line body test`}
	if !containsPost(posts, post) {
		t.Errorf("post not found, all posts %+v, search post %+v", posts, post)
	}
}

func containsPost(posts []templating.Post, post templating.Post) bool {
	for _, p := range posts {
		if reflect.DeepEqual(p, post) {
			return true
		}
	}
	return false
}
