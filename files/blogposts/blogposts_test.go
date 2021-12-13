package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/errantDev/blogposts"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Another failure")
}

// func TestBlogposts(t *testing.T) {
// 	t.Run("Correct number of posts created", func(t *testing.T) {
// 		fs := fstest.MapFS{
// 			"hello world.md":  {Data: []byte("hi")},
// 			"hello-world2.md": {Data: []byte("hola")},
// 		}

// 		posts, err := blogposts.NewPostsFromFS(fs)

// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if len(posts) != len(fs) {
// 			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 		}
// 	})
// 	t.Run("Error from file open is handled gracefully", func(t *testing.T) {
// 		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

// 		if err == nil {
// 			t.Error("Expected failure to be thrown")
// 		}
// 	})

// }

func assertPost(t *testing.T, got, want blogposts.Post) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestNewBlogposts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, testing
---
Another Post
Thanks`
	)

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}
