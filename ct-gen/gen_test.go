package mokuji

import "testing"

func TestGetMokuji(t *testing.T) {
	mdContent := []byte(`
# h1だよ
## h2だよ
### h3だよ
## h2だよ
### h3だよ
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
`)
	_, err := GetMokuji(mdContent)
	if err != nil {
		t.Fatal(err)
	}
}
