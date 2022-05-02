package blogpost

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	keyValueSeparator = ": "
	titleKey          = "Title"
	descriptionKey    = "Description"
	tagsKey           = "Tags"
)

func newPost(postFile io.Reader) (Post, error) {

	scanner := bufio.NewScanner(postFile)

	readInfoLine := func(key string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), key+keyValueSeparator)
	}

	title := readInfoLine(titleKey)
	desc := readInfoLine(descriptionKey)
	tags := strings.Split(readInfoLine(tagsKey), ", ")
	scanner.Scan() // Ignore "---" line
	body := readBody(scanner)

	post := Post{
		Title:       title,
		Description: desc,
		Tags:        tags,
		Body:        body,
	}

	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
