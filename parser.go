package mdparser

import (
	"strings"
)

type ParsedMD struct {
	Raw    string
	Blocks []BlockContent
}

type BlockContent struct {
	Contents []InlineContent
	Type     BlockContentType
}

type InlineContent struct {
	Text        string
	Type        InlineContentType
	Src         string
	Alt         string
	HasChildren bool
	Children    []InlineContent
}

func Parse(b []byte) (*ParsedMD, error) {
	md := ParsedMD{
		Raw: string(b),
	}

	lines := strings.Split(md.Raw, "\n")

	for _, line := range lines {
		for _, matcher := range blockContentMatchers() {
			if t, ok := matcher.match(line); ok {

				md.Blocks = append(md.Blocks, BlockContent{
					Type: t,
					Contents: []InlineContent{
						{
							Text: line,
							Type: InlineContentTypeText,
						},
					},
				})
			}
		}
	}

	return &md, nil
}
