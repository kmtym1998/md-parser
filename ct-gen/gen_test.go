package mokuji

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	parser "github.com/kmtym1998/md-parser"
	"github.com/stretchr/testify/assert"
)

func TestToNestableHeaderList(t *testing.T) {
	t.Run("1.md", func(t *testing.T) {
		mdContent, err := os.ReadFile("test/samples/1.md")
		if err != nil {
			t.Fatal(err)
		}

		md, err := parser.Parse(mdContent)
		if err != nil {
			t.Fatal(err)
		}

		nestedHeadingList, err := BlockList(md.Blocks).ToNestedHeaderList()
		if err != nil {
			t.Fatal(err)
		}

		assertToNestableHeaderList(
			t,
			nestedHeadingList,
			"test/samples/1.md",
			"test/golden_files/nested_heading_list_1.json",
		)
	})

	t.Run("2.md", func(t *testing.T) {
		mdContent, err := os.ReadFile("test/samples/2.md")
		if err != nil {
			t.Fatal(err)
		}

		md, err := parser.Parse(mdContent)
		if err != nil {
			t.Fatal(err)
		}

		nestedHeadingList, err := BlockList(md.Blocks).ToNestedHeaderList()
		if err != nil {
			t.Fatal(err)
		}

		assertToNestableHeaderList(
			t,
			nestedHeadingList,
			"test/samples/2.md",
			"test/golden_files/nested_heading_list_2.json",
		)
	})

	t.Run("3.md", func(t *testing.T) {
		mdContent, err := os.ReadFile("test/samples/3.md")
		if err != nil {
			t.Fatal(err)
		}

		md, err := parser.Parse(mdContent)
		if err != nil {
			t.Fatal(err)
		}

		nestedHeadingList, err := BlockList(md.Blocks).ToNestedHeaderList()
		if err != nil {
			t.Fatal(err)
		}

		assertToNestableHeaderList(
			t,
			nestedHeadingList,
			"test/samples/3.md",
			"test/golden_files/nested_heading_list_3.json",
		)
	})
}

func assertToNestableHeaderList(
	t *testing.T,
	nestedHeadingList NestableHeaderList,
	inputFilePath,
	goldenFilePath string,
) {
	t.Helper()

	actual, err := json.Marshal(nestedHeadingList)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFilePath)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("only header contents are included", func(t *testing.T) {
		for _, nestedHeading := range nestedHeadingList {
			assertOnlyHeaderIsIncluded(t, nestedHeading)
		}
	})

	t.Run("gets the same output as golden file", func(t *testing.T) {
		if !assert.JSONEq(t, string(expected), string(actual)) {
			var buf bytes.Buffer
			if err := json.Indent(&buf, actual, "", "  "); err != nil {
				t.Fatal(err)
			}

			if err := os.WriteFile(goldenFilePath, []byte(buf.String()+"\n"), os.ModePerm); err != nil {
				t.Fatal(err)
			}

			t.Log("actual output is overwritten to golden file")
		}
	})
}

func assertOnlyHeaderIsIncluded(t *testing.T, nestableHeader NestableHeader) {
	t.Helper()

	if !nestableHeader.Type.IsHeader() {
		t.Fatalf("unexpected type: %+v", nestableHeader.Type)
	}

	for _, child := range nestableHeader.Children {
		assertOnlyHeaderIsIncluded(t, child)
	}
}

func TestToContentTable(t *testing.T) {
	t.Run("1.md", func(t *testing.T) {
		nestedHeaderList := arrangeToContentTable(t, "test/samples/1.md")

		actual := nestedHeaderList.ToContentTable("  ")

		expected := `- h1 だよ
  - h3 だよ
    - h6 だよ
    - h6 だよ
  - h3 だよ
    - h5 だよ
    - h5 だよ
`

		assert.Equal(t, expected, actual)
	})

	t.Run("2.md", func(t *testing.T) {
		nestedHeaderList := arrangeToContentTable(t, "test/samples/2.md")

		actual := nestedHeaderList.ToContentTable("  ")

		expected := `- h1 だよ
  - h2 だよ
    - h3 だよ
  - h2 だよ
    - h3 だよ
      - h4 だよ
      - h4 だよ
      - h4 だよ
    - h3 だよ
      - h4 だよ
      - h4 だよ
    - h3 だよ
      - h4 だよ
        - h5 だよ
          - h6 だよ
    - h3 だよ
      - h4 だよ
      - h4 だよ
- h1 だよ
  - h2 だよ
    - h4 だよ
    - h4 だよ
`

		assert.Equal(t, expected, actual)
	})

	t.Run("3.md", func(t *testing.T) {
		nestedHeaderList := arrangeToContentTable(t, "test/samples/3.md")

		actual := nestedHeaderList.ToContentTable("  ")

		expected := `- はじめに
- 画像の圧縮処理をするコードサンプル
  - jpeg の圧縮
  - png の圧縮
- まとめ
- 参考
`

		assert.Equal(t, expected, actual)
	})
}

func arrangeToContentTable(t *testing.T, sampleMDFilePath string) NestableHeaderList {
	t.Helper()

	b, err := os.ReadFile(sampleMDFilePath)
	if err != nil {
		t.Fatal(err)
	}

	md, err := parser.Parse(b)
	if err != nil {
		t.Fatal(err)
	}

	nestedHeaderList, err := BlockList(md.Blocks).ToNestedHeaderList()
	if err != nil {
		t.Fatal(err)
	}

	return nestedHeaderList
}
