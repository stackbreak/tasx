package handlers

import (
	"errors"
	"io"
)

var ErrEmptyBodyJSON = errors.New("pkg.handlers: empty request body")

func isEmptyBodyErr(err error) error {
	if errors.Is(err, io.EOF) {
		err = ErrEmptyBodyJSON
	}
	return err
}
