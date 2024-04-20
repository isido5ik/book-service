package dtos

type Book struct {
	ID         int    `json:"-" db:"id"`
	Name       string `json:"name" db:"name"`
	Author     string `json:"author" db:"author"`
	PageNumber int    `json:"pagenumber" db:"page_number"`
	Date       int    `json:"date" db:"date"`
	Rating     int    `json:"rating" db:"rating"`
}
