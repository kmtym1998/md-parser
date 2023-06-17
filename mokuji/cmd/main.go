package main

import (
	"fmt"
	"os"

	parser "github.com/kmtym1998/md-parser"
	"github.com/kmtym1998/md-parser/mokuji"
)

func main() {
	mdContent, _ := os.ReadFile("test/samples/3.md")

	md, err := parser.Parse(mdContent)
	if err != nil {
		panic(err)
	}

	nestedHeaderList, err := mokuji.BlockList(md.Blocks).ToNestedHeaderList()
	if err != nil {
		panic(err)
	}

	fmt.Println(nestedHeaderList.ToContentTable("  "))
}
