package blogposts

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
	TitleSeparator       = "Title: "
	DescriptionSeparator = "Description: "
	TagsSeparator        = "Tags: "
	BodySeparator        = "Body: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(separator string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), separator)
	}

	readBody := func() string {
		scanner.Scan()
		buf := bytes.Buffer{}
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		return strings.TrimSuffix(buf.String(), "\n")
	}

	return Post{
		Title:       readMetaLine(TitleSeparator),
		Description: readMetaLine(DescriptionSeparator),
		Tags:        strings.Split(readMetaLine(TagsSeparator), ", "),
		Body:        readBody(),
	}, nil
}
