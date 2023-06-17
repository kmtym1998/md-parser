package mokuji

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	parser "github.com/kmtym1998/md-parser"
	"github.com/samber/lo"
)

type BlockList []parser.BlockContent

type NestableHeadingList []NestableHeading

type NestableHeading struct {
	Text     string
	Type     parser.BlockContentType
	Children NestableHeadingList
}

func GetMokuji(mdContent []byte) (string, error) {
	md, err := parser.Parse(mdContent)
	if err != nil {
		return "", err
	}

	headingBlocks := lo.Filter(md.Blocks, func(b parser.BlockContent, _ int) bool {
		return lo.Contains([]parser.BlockContentType{
			parser.BlockContentTypeHeader1,
			parser.BlockContentTypeHeader2,
			parser.BlockContentTypeHeader3,
			parser.BlockContentTypeHeader4,
			parser.BlockContentTypeHeader5,
			parser.BlockContentTypeHeader6,
		}, b.Type)
	})

	nestedHeadingList, err := getNestableHeadingList(headingBlocks)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(nestedHeadingList)
	if err != nil {
		return "", err
	}

	f, _ := os.Create("hoge.json")
	fmt.Fprint(f, string(b))

	return "", nil
}

func getNestableHeadingList(headingBlocks BlockList) (nestableHeadingList NestableHeadingList, err error) {
	for _, hb := range headingBlocks {
		if len(hb.Contents) == 0 {
			return nil, errors.New("heading block has no contents")
		}

		if len(hb.Contents[0].ContainedTypes) == 0 {
			return nil, errors.New("heading block has no contained types")
		}

		nestableHeadingList = nestableHeadingList.append(hb)
	}

	return nestableHeadingList, nil
}

func (l NestableHeadingList) append(block parser.BlockContent) NestableHeadingList {
	if len(l) == 0 {
		l = append(l, NestableHeading{
			Text: block.Contents[0].Text,
			Type: block.Type,
		})

		return l
	}

	lastParentHeading := &l[len(l)-1]

	if block.Type <= lastParentHeading.Type {
		return append(l, NestableHeading{
			Text: block.Contents[0].Text,
			Type: block.Type,
		})
	}

	lastParentHeading.Children = lastParentHeading.Children.append(block)

	return l
}
