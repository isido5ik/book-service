package handler

import (
	"RedisLesson/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	books := router.Group("/books")
	{
		books.POST("/", h.createBook)
		books.GET("/", h.getAllBooks)
		books.GET("/:id", h.getBookById)
		books.PUT("/:id", h.updateBook)
		books.DELETE(":id", h.deleteBook)
	}
	return router
}
