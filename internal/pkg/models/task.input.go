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

type InputUpdateTask struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsDone      *bool   `json:"is_done"`
}

func (i *InputUpdateTask) Validate() error {
	if i.Title == nil && i.Description == nil && i.IsDone == nil {
		return ErrEmptyInput
	}

	return nil
}
