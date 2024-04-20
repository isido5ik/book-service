package dtos

import "errors"

type BookFilter struct {
	Name          string `json:"name" db:"name"`
	Author        string `json:"author" db:"author"`
	PageNumberMin int    `json:"pageNumberMin" db:"-"`
	PageNumberMax int    `json:"pageNumberMax" db:"-"`
	Date          int    `json:"date" db:"date"`
	Rating        int    `json:"rating" db:"rating"`
}

func (f BookFilter) Validate() error {
	if f.PageNumberMax != 0 {
		if f.PageNumberMin > f.PageNumberMax {
			return errors.New("max < min")
		}
	}

	return nil
}
