package mdparser

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
)

type InlineContentType string

const (
	InlineContentTypeBold      InlineContentType = "bold"
	InlineContentTypeItalic    InlineContentType = "italic"
	InlineContentTypeUnderline InlineContentType = "underline"
	InlineContentTypeStrike    InlineContentType = "strike"
	InlineContentTypeLink      InlineContentType = "link"
	InlineContentTypeImage     InlineContentType = "image"
	InlineContentTypeCode      InlineContentType = "code"
	InlineContentTypeText      InlineContentType = "text"
)

func Parse(b []byte) (ParsedMD, error) {
	md := ParsedMD{
		Raw: string(b),
	}

	return md, nil
}
