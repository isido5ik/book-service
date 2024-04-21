package dtos

type Book struct {
	ID         int    `json:"-" db:"id"`
	Name       string `json:"name" db:"name" binding:"required"`
	Author     string `json:"author" db:"author" binding:"required"`
	PageNumber int    `json:"pagenumber" db:"page_number" binding:"required"`
	Date       int    `json:"date" db:"date" binding:"required"`
	Rating     int    `json:"rating" db:"rating" binding:"required"`
}
