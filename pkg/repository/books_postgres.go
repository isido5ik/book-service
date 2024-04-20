package repository

import (
	"RedisLesson/dtos"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// type BooksPostgres struct {
// 	db    *sqlx.DB
// 	redis *redis.Client
// }

// func NewBooksPostgres(db *sqlx.DB) *BooksPostgres {
// 	return &BooksPostgres{
// 		db:    db,
// 		redis: NewRedisClient(),
// 	}
// }

func (r *repository) CreateBook(book dtos.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, author, page_number, date, rating) VALUES($1, $2, $3, $4, $5) RETURNING id", booksTable)
	row := r.db.QueryRow(query, book.Name, book.Author, book.PageNumber, book.Date, book.Rating)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetAllBooks() ([]dtos.Book, error) {

	var books []dtos.Book
	query := fmt.Sprintf("SELECT * FROM %s ", booksTable)
	err := r.db.Select(&books, query)

	return books, err
}

// func (r *repository) GetBookById(id int) (dtos.Book, error) {

// 	var book dtos.Book

// 	JSONbook, redisErr := r.redis.Get(context.Background(), strconv.Itoa(id)).Result()
// 	if redisErr == nil {
// 		json.Unmarshal([]byte(JSONbook), &book)
// 		log.Println("Getting element from redis")
// 		return book, nil
// 	}

// 	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", booksTable)
// 	err := r.db.Get(&book, query, id)

// 	marshalledBook, marshalErr := json.Marshal(book)
// 	if marshalErr != nil {
// 		return book, marshalErr
// 	}

// 	redisErr = r.redis.Set(context.Background(), strconv.Itoa(book.ID), marshalledBook, 0).Err()
// 	if redisErr != nil {
// 		return book, redisErr
// 	}
// 	return book, err
// }

// func (r *repository) DeleteBook(id int) error {
// 	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTable)
// 	_, err := r.db.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *repository) GetBookById(id int) (dtos.Book, error) {

	book, err := r.getBookFromRedis(id)
	if err == nil {
		log.Println("Getting element from Redis")
		return book, nil
	}

	book, err = r.getBookFromDB(id)
	if err != nil {
		return dtos.Book{}, err
	}

	// Кэшируем книгу в Redis
	if err := r.cacheBookToRedis(book); err != nil {
		log.Println("Error caching book to Redis:", err)
	}

	return book, nil
}

func (r *repository) DeleteBook(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	if err := r.deleteBookFromRedis(id); err != nil {
		log.Println("Error deleting book from Redis:", err)
	}
	return nil
}

func (r *repository) getBookFromRedis(id int) (dtos.Book, error) {
	var book dtos.Book
	JSONbook, err := r.redis.Get(context.Background(), strconv.Itoa(id)).Result()
	if err != nil {
		return dtos.Book{}, err
	}
	if err := json.Unmarshal([]byte(JSONbook), &book); err != nil {
		return dtos.Book{}, err
	}
	return book, nil
}

func (r *repository) cacheBookToRedis(book dtos.Book) error {
	marshalledBook, err := json.Marshal(book)
	if err != nil {
		return err
	}
	if err := r.redis.Set(context.Background(), strconv.Itoa(book.ID), marshalledBook, time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

func (r *repository) deleteBookFromRedis(id int) error {
	if err := r.redis.Del(context.Background(), strconv.Itoa(id)).Err(); err != nil {
		return err
	}
	return nil
}

func (r *repository) getBookFromDB(id int) (dtos.Book, error) {
	var book dtos.Book
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", booksTable)
	err := r.db.Get(&book, query, id)
	return book, err
}

//_____________________________________________________________________________________________

func (r *repository) UpdateBook(id int, input dtos.BookUpdateInput) error {
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
