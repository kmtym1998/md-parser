package mdparser

import "strings"

type BlockContentType int

const (
	BlockContentTypeHeader1     BlockContentType = iota
	BlockContentTypeHeader2     BlockContentType = iota
	BlockContentTypeHeader3     BlockContentType = iota
	BlockContentTypeHeader4     BlockContentType = iota
	BlockContentTypeHeader5     BlockContentType = iota
	BlockContentTypeHeader6     BlockContentType = iota
	BlockContentTypeQuote       BlockContentType = iota
	BlockContentTypeCode        BlockContentType = iota
	BlockContentTypeList        BlockContentType = iota
	BlockContentTypeOrderedList BlockContentType = iota
	BlockContentTypeImage       BlockContentType = iota
	BlockContentTypeLink        BlockContentType = iota
	BlockContentTypeTable       BlockContentType = iota
	BlockContentTypeHorizontal  BlockContentType = iota
	BlockContentTypeParagraph   BlockContentType = iota
	BlockContentTypeEmpty       BlockContentType = iota
	BlockContentTypeUnknown     BlockContentType = iota
)

type InlineContentType int

const (
	InlineContentTypeBold      InlineContentType = iota
	InlineContentTypeItalic    InlineContentType = iota
	InlineContentTypeUnderline InlineContentType = iota
	InlineContentTypeStrike    InlineContentType = iota
	InlineContentTypeLink      InlineContentType = iota
	InlineContentTypeImage     InlineContentType = iota
	InlineContentTypeCode      InlineContentType = iota
	InlineContentTypeText      InlineContentType = iota
)

type blockContentMatcher interface {
	match(content string) (BlockContentType, bool)
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

func (m header1Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader1, true
	}

	return BlockContentTypeUnknown, false
}

type header2Matcher string

func (m header2Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader2, true
	}

	return BlockContentTypeUnknown, false
}

type header3Matcher string

func (m header3Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader3, true
	}

	return BlockContentTypeUnknown, false
}

type header4Matcher string

func (m header4Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader4, true
	}

	return BlockContentTypeUnknown, false
}

type header5Matcher string

func (m header5Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader5, true
	}

	return BlockContentTypeUnknown, false
}

type header6Matcher string

func (m header6Matcher) match(content string) (BlockContentType, bool) {
	if strings.HasPrefix(content, string(m)) {
		return BlockContentTypeHeader6, true
	}

	return BlockContentTypeUnknown, false
}

type emptyMatcher string

func (m emptyMatcher) match(content string) (BlockContentType, bool) {
	if content == string(m) {
		return BlockContentTypeEmpty, true
	}

	return BlockContentTypeUnknown, false
}
