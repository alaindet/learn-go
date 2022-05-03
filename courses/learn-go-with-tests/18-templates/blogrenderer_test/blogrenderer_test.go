package blogrenderer_test

import (
	"bytes"
	"learn_go_with_tests/templates/blogrenderer"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {

	var aPost = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		result := buf.String()
		expected := joinStrings(
			"<h1>hello world</h1>",
			"<p>This is a description</p>",
			"Tags: <ul><li>go</li><li>tdd</li></ul>",
		)

		if result != expected {
			t.Errorf("got '%s' want '%s'", result, expected)
		}
	})
}

// func joinLines(lines ...string) string {
// 	return strings.Join(lines, "\n")
// }

func joinStrings(lines ...string) string {
	return strings.Join(lines, "")
}
