package service

import (
	"RedisLesson/dtos"
	"RedisLesson/pkg/repository"
)

type BookService struct {
	repo repository.Books
}

func NewBookService(repo repository.Books) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Create(book dtos.Book) (int, error) {
	return s.repo.Create(book)
}

func (s *BookService) GetAll() ([]dtos.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetById(id int) (dtos.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *BookService) Update(id int, input dtos.BookUpdateInput) error {
	return s.repo.Update(id, input)
}
