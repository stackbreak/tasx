package handlers

import (
	"errors"
	"io"
)

var (
	ErrEmptyBodyJSON        = errors.New("pkg.handlers: empty request body")
	ErrInvalidListIdParam   = errors.New("pkg.handlers: invalid list_id param")
	ErrInvalidTaskIdParam   = errors.New("pkg.handlers: invalid task_id param")
	ErrPersonIdNotExtracted = errors.New("pkg.handlers: unable to extract person_id")
)

func isEmptyBodyErr(err error) error {
	if errors.Is(err, io.EOF) {
		err = ErrEmptyBodyJSON
	}
	return err
}
