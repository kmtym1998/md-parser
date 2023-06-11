package mdparser

import "strings"

type BlockContentType string

const (
	BlockContentTypeHeader1     BlockContentType = "header1"
	BlockContentTypeHeader2     BlockContentType = "header2"
	BlockContentTypeHeader3     BlockContentType = "header3"
	BlockContentTypeHeader4     BlockContentType = "header4"
	BlockContentTypeHeader5     BlockContentType = "header5"
	BlockContentTypeHeader6     BlockContentType = "header6"
	BlockContentTypeQuote       BlockContentType = "quote"
	BlockContentTypeCode        BlockContentType = "code"
	BlockContentTypeList        BlockContentType = "list"
	BlockContentTypeOrderedList BlockContentType = "orderedList"
	BlockContentTypeImage       BlockContentType = "image"
	BlockContentTypeLink        BlockContentType = "link"
	BlockContentTypeTable       BlockContentType = "table"
	BlockContentTypeHorizontal  BlockContentType = "horizontal"
	BlockContentTypeParagraph   BlockContentType = "paragraph"
	BlockContentTypeEmpty       BlockContentType = "empty"
	BlockContentTypeUnknown     BlockContentType = "unknown"
)

type blockContentMatcher interface {
	match(content string) (BlockContentType, bool)
	trimText(content string) string
}

func blockContentMatchers() []blockContentMatcher {
	return []blockContentMatcher{
		header1Matcher("# "),
		header2Matcher("## "),
		header3Matcher("### "),
		header4Matcher("#### "),
		header5Matcher("##### "),
		header6Matcher("###### "),
		emptyMatcher(""),
	}
}

type header1Matcher string

func (m header1Matcher) trimText(content string) string {
	return strings.TrimPrefix(content, string(m))
}

func (m header1Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader1, true
	}

	return BlockContentTypeUnknown, false
}

type header2Matcher string

func (m header2Matcher) trimText(content string) string {
	return strings.TrimPrefix(content, string(m))
}

func (m header2Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader2, true
	}

	return BlockContentTypeUnknown, false
}

type header3Matcher string

func (m header3Matcher) trimText(content string) string {
	return strings.TrimPrefix(content, string(m))
}

func (m header3Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader3, true
	}

	return BlockContentTypeUnknown, false
}

type header4Matcher string

func (m header4Matcher) trimText(content string) string {
	return strings.TrimPrefix(content, string(m))
}

func (m header4Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader4, true
	}

	return BlockContentTypeUnknown, false
}

type header5Matcher string

func (m header5Matcher) trimText(content string) string {
	return strings.TrimPrefix(content, string(m))
}

func (m header5Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader5, true
	}

	return BlockContentTypeUnknown, false
}

type header6Matcher string

func (m header6Matcher) trimText(content string) string {
	return strings.TrimPrefix(content, string(m))
}

func (m header6Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader6, true
	}

	return BlockContentTypeUnknown, false
}

type emptyMatcher string

func (m emptyMatcher) trimText(content string) string {
	return ""
}

func (m emptyMatcher) match(content string) (BlockContentType, bool) {
	if content == string(m) {
		return BlockContentTypeEmpty, true
	}

	return BlockContentTypeUnknown, false
}
