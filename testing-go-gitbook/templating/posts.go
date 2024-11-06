package templating

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func ReadPostsFromFS(fileSystem fs.FS, dirName string) ([]Post, error) {
	var posts []Post

	dirEntry, err := fs.ReadDir(fileSystem, dirName)
	if err != nil {
		return nil, err
	}
	for _, de := range dirEntry {
		if de.IsDir() {
			continue
		}

		filePath := filepath.Join(dirName, de.Name())
		file, err := fileSystem.Open(filePath)
		if err != nil {
			continue
		}

		post := newPost(file)
		posts = append(posts, post)

		file.Close()
	}

	return posts, nil
}

func newPost(r io.Reader) Post {
	getPostData := postDataReader(r)

	title := getPostData("Title: ")
	description := getPostData("Description: ")
	tags := strings.Split(getPostData("Tags: "), ", ")
	body := getPostData("#body#")

	return Post{title, description, tags, body}
}

func postDataReader(r io.Reader) func(string) string {
	scanner := bufio.NewScanner(r)

	return func(prefix string) string {
		scanner.Scan()

		if prefix == "#body#" {
			var b bytes.Buffer
			for scanner.Scan() {
				fmt.Fprintln(&b, scanner.Text())
			}
			return strings.TrimSuffix(b.String(), "\n")
		}

		return strings.TrimPrefix(scanner.Text(), prefix)
	}
}
