package mokuji

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	parser "github.com/kmtym1998/md-parser"
	"github.com/stretchr/testify/assert"
)

const mdStr = `# h1だよ
## h2だよ

aaaaaa
aaaa

### h3だよ
## h2だよ
### h3だよ

bbbb
bbbbb
bbbbbbb # bbb

#### h4だよ
#### h4だよ
#### h4だよ
### h3だよ
#### h4だよ
#### h4だよ
### h3だよ
#### h4だよ
##### h5だよ
###### h6だよ
### h3だよ
#### h4だよ
#### h4だよ
# h1だよ
## h2だよ
#### h4だよ
#### h4だよ
`

func TestGetMokuji(t *testing.T) {
	mdContent := []byte(mdStr)
	_, err := GetMokuji(mdContent)
	if err != nil {
		t.Fatal(err)
	}
}

func TestToNestableHeaderList(t *testing.T) {
	mdContent := []byte(mdStr)
	md, err := parser.Parse(mdContent)
	if err != nil {
		t.Fatal(err)
	}

	nestedHeadingList, err := BlockList(md.Blocks).toNestedHeaderList()
	if err != nil {
		t.Fatal(err)
	}

	actual, err := json.Marshal(nestedHeadingList)
	if err != nil {
		t.Fatal(err)
	}

	const goldenFilePath = "test/golden_files/nested_heading_list.json"
	expected, err := os.ReadFile(goldenFilePath)
	if err != nil {
		t.Fatal(err)
	}

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
}
