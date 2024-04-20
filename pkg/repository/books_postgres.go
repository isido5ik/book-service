package repository

import (
	"RedisLesson/dtos"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type BooksPostgres struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewBooksPostgres(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{
		db:    db,
		redis: NewRedisClient(),
	}
}

func (r *BooksPostgres) Create(book dtos.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, author, page_number, date, rating) VALUES($1, $2, $3, $4, $5) RETURNING id", booksTable)
	row := r.db.QueryRow(query, book.Name, book.Author, book.PageNumber, book.Date, book.Rating)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BooksPostgres) GetAll() ([]dtos.Book, error) {

	var books []dtos.Book
	query := fmt.Sprintf("SELECT * FROM %s ", booksTable)
	err := r.db.Select(&books, query)

	return books, err
}

func (r *BooksPostgres) GetById(id int) (dtos.Book, error) {

	var book dtos.Book

	JSONbook, redisErr := r.redis.Get(context.Background(), strconv.Itoa(id)).Result()
	if redisErr == nil {
		json.Unmarshal([]byte(JSONbook), &book)
		log.Println("Getting element from redis")
		return book, nil
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", booksTable)
	err := r.db.Get(&book, query, id)

	marshalledBook, marshalErr := json.Marshal(book)
	if marshalErr != nil {
		return book, marshalErr
	}

	redisErr = r.redis.Set(context.Background(), strconv.Itoa(book.ID), marshalledBook, 0).Err()
	if redisErr != nil {
		return book, redisErr
	}
	return book, err
}

func (r *BooksPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *BooksPostgres) Update(id int, input dtos.BookUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, input.Author)
		argId++
	}

	if input.PageNumber != nil {
		setValues = append(setValues, fmt.Sprintf("page_number=$%d", argId))
		args = append(args, input.PageNumber)
		argId++
	}

	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, input.Date)
		argId++
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, input.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", booksTable, setQuery, id)

	logrus.Printf("making query for update: %s/n with args: %s", setQuery, args)
	_, err := r.db.Exec(query, args...)
	return err
}
