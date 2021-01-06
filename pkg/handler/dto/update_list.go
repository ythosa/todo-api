package dto

import "errors"

type UpdateList struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (ul *UpdateList) Validate() error {
	if ul.Title == nil && ul.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
