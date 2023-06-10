package mdparser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	b, err := os.ReadFile("samples/01.md")
	if err != nil {
		t.Fatal(err)
	}

	md, err := Parse(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(b), md.Raw)
}
