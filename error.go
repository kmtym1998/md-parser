package mdparser

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

type ParseError struct {
	content string
	msg     string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", e.msg)
}

func (e ParseError) Content() string {
	return e.content
}

func (e ParseError) Unwrap() error {
	return errors.Unwrap(e)
}

func NewParseError(content, msg string) error {
	return errors.WithStack(ParseError{
		content: content,
		msg:     msg,
	})
}
