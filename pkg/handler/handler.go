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
		books.POST("/create", h.createBook)
		books.GET("/get-all", h.getAllBooks)
		books.POST("/get-by-filter", h.getBooksByFilter)
		books.GET("/get/:id", h.getBookById)
		books.PUT("/put/:id", h.updateBook)
		books.DELETE("/delete/:id", h.deleteBook)
	}
	return router
}
