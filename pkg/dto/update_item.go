package dto

import "errors"

type UpdateItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *string `json:"done"`
}

func (ui *UpdateItem) Validate() error {
	if ui.Title == nil && ui.Description == nil && ui.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
