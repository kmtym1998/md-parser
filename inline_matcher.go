package mdparser

type InlineContentType string

func (t InlineContentType) String() string {
	return string(t)
}

const (
	InlineContentTypeBold      InlineContentType = "bold"
	InlineContentTypeItalic    InlineContentType = "italic"
	InlineContentTypeUnderline InlineContentType = "underline"
	InlineContentTypeStrike    InlineContentType = "strike"
	InlineContentTypeLink      InlineContentType = "link"
	InlineContentTypeImage     InlineContentType = "image"
	InlineContentTypeCode      InlineContentType = "code"
	InlineContentTypeText      InlineContentType = "text"
	InlineContentTypeUnknown   InlineContentType = "unknown"
)

type inlineContentMatcher interface {
	match(content string) *InlineSig
}

type InlineSig struct {
	Start string
	End   string
	Type  InlineContentType
}

func inlineContentMatchers() []inlineContentMatcher {
	return []inlineContentMatcher{
		boldMatcher(InlineSig{}),
	}
}

type boldMatcher InlineSig

func (m boldMatcher) match(content string) *InlineSig {
	return nil
}
