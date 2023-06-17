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
	Level    int
	Text     string
	Type     parser.BlockContentType
	Children []NestableHeading
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

		if len(nestableHeadingList) == 0 {
			nestableHeadingList = append(nestableHeadingList, NestableHeading{
				Text: hb.Contents[0].Text,
				Type: hb.Type,
			})
			continue
		}

		lastHeading := nestableHeadingList[len(nestableHeadingList)-1]

		if hb.Type < lastHeading.Type {
			nestableHeadingList[len(nestableHeadingList)-1].Children = append(lastHeading.Children, NestableHeading{
				Text: hb.Contents[0].Text,
				Type: hb.Type,
			})
			continue
		}

		nestableHeadingList = append(nestableHeadingList, NestableHeading{
			Text: hb.Contents[0].Text,
			Type: hb.Type,
		})
	}

	return nestableHeadingList, nil
}
