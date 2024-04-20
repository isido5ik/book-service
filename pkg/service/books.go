package service

import (
	"RedisLesson/dtos"
)

// type BookService struct {
// 	repo repository.Repository
// }

// func NewBookService(repo repository.Repository) *BookService {
// 	return &BookService{repo: repo}
// }

func (s *service) CreateBook(book dtos.Book) (int, error) {
	return s.repo.CreateBook(book)
}

func (s *service) GetAllBooks() ([]dtos.Book, error) {
	//try_from_redis

	//if error try from data base

	//or just write in one function
	return s.repo.GetAllBooks()
}

func (s *service) GetBooksByFilter(filter dtos.BookFilter) ([]dtos.Book, error) {
	return s.repo.GetBooksByFilter(filter)
}

func (s *service) GetBookById(id int) (dtos.Book, error) {
	return s.repo.GetBookById(id)
}

func (s *service) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}

func (s *service) UpdateBook(id int, input dtos.BookUpdateInput) error {
	return s.repo.UpdateBook(id, input)
}
