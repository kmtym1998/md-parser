package mokuji

import (
	"errors"
	"fmt"
	"strings"

	parser "github.com/kmtym1998/md-parser"
)

type BlockList []parser.BlockContent

type NestableHeaderList []NestableHeader

type NestableHeader struct {
	Depth    int
	Text     string
	Type     parser.BlockContentType
	Children NestableHeaderList
}

func (l BlockList) ToNestedHeaderList() (nestableHeaderList NestableHeaderList, err error) {
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

		nestableHeaderList = nestableHeaderList.append(block, 0)
	}

	return nestableHeaderList, nil
}

func (l NestableHeaderList) append(block parser.BlockContent, depth int) NestableHeaderList {
	if len(l) == 0 {
		l = append(l, NestableHeader{
			Text:  block.Contents[0].Text,
			Type:  block.Type,
			Depth: depth,
		})

		return l
	}

	lastParentHeader := &l[len(l)-1]

	if block.Type <= lastParentHeader.Type {
		return append(l, NestableHeader{
			Text:  block.Contents[0].Text,
			Type:  block.Type,
			Depth: lastParentHeader.Depth,
		})
	}

	lastParentHeader.Children = lastParentHeader.Children.append(block, depth+1)

	return l
}

func (l NestableHeaderList) ToContentTable(indent string) (result string) {
	for _, block := range l {
		result += fmt.Sprintf(
			"%s- %s\n",
			strings.Repeat(indent, block.Depth),
			block.Text,
		)

		for _, child := range block.Children {
			result += fmt.Sprintf(
				"%s- %s\n",
				strings.Repeat(indent, child.Depth),
				child.Text,
			)

			result += child.Children.ToContentTable(indent)
		}
	}

	return
}
