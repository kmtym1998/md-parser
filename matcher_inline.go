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
	match(content string) InlineContentType
	trimText(content string) string
	trimSrc(content string) string
	trimAlt(content string) string
}

type InlineSig struct {
	regExp string
	Type   InlineContentType
}

func inlineContentMatchers() []inlineContentMatcher {
	return []inlineContentMatcher{
		boldMatcher{regExp: `^\*\*.*\*\*$`, Type: InlineContentTypeBold},
	}
}

type boldMatcher InlineSig

func (m boldMatcher) match(content string) InlineContentType {
	return "unknown"
}
func (m boldMatcher) trimText(content string) string {
	return ""
}
func (m boldMatcher) trimSrc(content string) string {
	return ""
}
func (m boldMatcher) trimAlt(content string) string {
	return ""
}
