package blogpost_test

import (
	"errors"
	"io/fs"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"

	"learn_go_with_tests/reading_files/blogpost"
)

func joinLines(lines ...string) string {
	return strings.Join(lines, "\n")
}

func TestNewBlogPosts(t *testing.T) {

	fs := fstest.MapFS{
		"hello_world.md": {Data: []byte(joinLines(
			"Title: Post 1",
			"Description: Description 1",
			"Tags: tdd, go",
			"---",
			"This is the content of hello_world.md",
		))},
		"hello_world2.md": {Data: []byte(joinLines(
			"Title: Post 2",
			"Description: Description 2",
			"Tags: rust, borrow-checker",
			"---",
			"This is the content of hello_world2.md",
		))},
	}

	posts, err := blogpost.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	// Check first post
	assertPost(t, posts[0], blogpost.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        "This is the content of hello_world.md",
	})
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("This always fails")
}

func assertPost(t *testing.T, got blogpost.Post, want blogpost.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
