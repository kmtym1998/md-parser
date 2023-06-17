package mokuji

import (
	"encoding/json"
	"errors"
	"strings"

	parser "github.com/kmtym1998/md-parser"
)

type BlockList []parser.BlockContent

type NestableHeaderList []NestableHeader

type NestableHeader struct {
	Text     string
	Type     parser.BlockContentType
	Children NestableHeaderList
}

func GetMokuji(mdContent []byte) (string, error) {
	md, err := parser.Parse(mdContent)
	if err != nil {
		return "", err
	}

	nestedHeaderList, err := BlockList(md.Blocks).toNestedHeaderList()
	if err != nil {
		return "", err
	}

	json.Marshal(nestedHeaderList)

	return "", nil
}

func (l BlockList) toNestedHeaderList() (nestableHeaderList NestableHeaderList, err error) {
	for _, block := range l {
		if !strings.HasPrefix(block.Type.String(), "header") {
			continue
		}

		if len(block.Contents) == 0 {
			return nil, errors.New("header block has no contents")
		}

		if len(block.Contents[0].ContainedTypes) == 0 {
			return nil, errors.New("header block has no contained types")
		}

		nestableHeaderList = nestableHeaderList.append(block)
	}

	return nestableHeaderList, nil
}

func (l NestableHeaderList) append(block parser.BlockContent) NestableHeaderList {
	if len(l) == 0 {
		l = append(l, NestableHeader{
			Text: block.Contents[0].Text,
			Type: block.Type,
		})

		return l
	}

	lastParentHeader := &l[len(l)-1]

	if block.Type <= lastParentHeader.Type {
		return append(l, NestableHeader{
			Text: block.Contents[0].Text,
			Type: block.Type,
		})
	}

	lastParentHeader.Children = lastParentHeader.Children.append(block)

	return l
}
