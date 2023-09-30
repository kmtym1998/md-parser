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
	// FIXME: Deprecated
	ContainedTypes []InlineContentType // deprecated
	Text           string
	Src            string
	Alt            string
	Children       []InlineContent
}

func Parse(b []byte) (*ParsedMD, error) {
	md := ParsedMD{
		Raw: string(b),
	}

	lines := strings.Split(md.Raw, "\n")

	for _, line := range lines {
		for _, matcher := range blockContentMatchers() {
			if blockType, ok := matcher.match(line); ok {
				md.Blocks = append(md.Blocks, BlockContent{
					Type: blockType,
					Contents: []InlineContent{{
						Text:     matcher.trimText(line),
						Children: parseInlineContent(matcher.trimText(line)),
					}},
				})

				continue
			}
		}
	}

	return &md, nil
}

func parseInlineContent(line string) []InlineContent {
	var contents []InlineContent

	for _, matcher := range inlineContentMatchers() {
		if inlineType := matcher.match(line); inlineType != InlineContentTypeUnknown {
			txt := matcher.trimText(line)
			contents = append(contents, InlineContent{
				ContainedTypes: []InlineContentType{InlineContentTypeText},
				Text:           txt,
				Src:            matcher.trimSrc(line),
				Alt:            matcher.trimAlt(line),
				Children:       parseInlineContent(txt),
			})
		}
	}

	return contents
}
