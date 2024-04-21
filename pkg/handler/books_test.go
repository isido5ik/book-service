package handler

import (
	"RedisLesson/dtos"
	"RedisLesson/pkg/service"
	mock_service "RedisLesson/pkg/service/mock"
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_createBook(t *testing.T) {
	type mockBehavior func(s *mock_service.MockService, book dtos.Book)

	testTable := []struct {
		name                string
		inputBody           string
		inputBook           dtos.Book
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "author":"test", "pagenumber":100, "date":2024, "rating":1}`,
			inputBook: dtos.Book{
				Name:       "Test",
				Author:     "test",
				PageNumber: 100,
				Date:       2024,
				Rating:     1,
			},
			mockBehavior: func(s *mock_service.MockService, book dtos.Book) {
				s.EXPECT().CreateBook(book).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Empty Fields",
			inputBody: `{"author":"test", "pagenumber":100, "date":2024, "rating":1}`,
			mockBehavior: func(s *mock_service.MockService, book dtos.Book) {

			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},

		{
			name:      "Service Failure",
			inputBody: `{"name":"Test", "author":"test", "pagenumber":100, "date":2024, "rating":1}`,
			inputBook: dtos.Book{
				Name:       "Test",
				Author:     "test",
				PageNumber: 100,
				Date:       2024,
				Rating:     1,
			},
			mockBehavior: func(s *mock_service.MockService, book dtos.Book) {
				s.EXPECT().CreateBook(book).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			books := mock_service.NewMockService(c)
			testCase.mockBehavior(books, testCase.inputBook)

			services := service.NewService(books)
			handler := NewHandler(services)

			//test Server
			r := gin.New()
			r.POST("/create", handler.createBook)

			//test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create", bytes.NewBufferString(testCase.inputBody))

			//perform request
			r.ServeHTTP(w, req)

			//assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_getAllBooks(t *testing.T) {
	type mockBehavior func(s *mock_service.MockService)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockService) {
				fakeBooks := []dtos.Book{
					{ID: 1, Name: "Book 1", Author: "Author 1", PageNumber: 100, Date: 2022, Rating: 5},
					{ID: 2, Name: "Book 2", Author: "Author 2", PageNumber: 200, Date: 2023, Rating: 4},
				}
				s.EXPECT().GetAllBooks(dtos.PaginationParams{Page: 1, PageSize: 2}).Return(fakeBooks, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"name":"Book 1","author":"Author 1","pagenumber":100,"date":2022,"rating":5},{"name":"Book 2","author":"Author 2","pagenumber":200,"date":2023,"rating":4}],"page":1,"pageSize":2}`,
		},
		{
			name: "Service Failure",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().GetAllBooks(dtos.PaginationParams{Page: 1, PageSize: 2}).Return(nil, errors.New("service error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			books := mock_service.NewMockService(c)
			testCase.mockBehavior(books)

			services := service.NewService(books)
			handler := NewHandler(services)

			// Инициализация фейкового контекста Gin
			r := gin.New()
			r.GET("/books", handler.getAllBooks)

			// Выполнение запроса
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/books?page=1&pageSize=2", nil)
			r.ServeHTTP(w, req)

			// Проверка статуса кода ответа
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

			// Проверка тела ответа
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
