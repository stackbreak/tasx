package models

import "errors"

var ErrEmptyInput = errors.New("pkg.models: empty input")

type InputUpdateTaskList struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i *InputUpdateTaskList) Validate() error {
	if i.Title == nil && i.Description == nil {
		return ErrEmptyInput
	}

	return nil
}
