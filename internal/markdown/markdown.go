package markdown

import (
	"bytes"

	obsidian "github.com/powerman/goldmark-obsidian"
	"github.com/yuin/goldmark"
)

func Markdown2Html(markdown string) string {
	md := goldmark.New(
		goldmark.WithExtensions(
			obsidian.NewObsidian(),
		),
	)

	var buf bytes.Buffer

	if err := md.Convert([]byte(markdown), &buf); err != nil {
		panic(err)
	}

	return buf.String()
}
