package dtos

import "errors"

type BookUpdateInput struct {
	Name       *string `json:"name" db:"name"`
	Author     *string `json:"author" db:"author"`
	PageNumber *int    `json:"pagenumber" db:"page_number"`
	Date       *int    `json:"date" db:"date"`
	Rating     *int    `json:"rating" db:"rating"`
}

func (i BookUpdateInput) Validate() error {
	if i.Name == nil && i.Author == nil && i.PageNumber == nil && i.Date == nil && i.Rating == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
