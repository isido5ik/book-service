package handler

import (
	"RedisLesson/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createBook(c *gin.Context) {
	var input dtos.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err := h.services.CreateBook(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllBooksResponse struct {
	Data     []dtos.Book `json:"data"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

func (h *Handler) getAllBooks(c *gin.Context) {
	var pagination dtos.PaginationParams

	var err error
	pagination.Page, pagination.PageSize, err = dtos.ValidatePage(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	var books []dtos.Book
	books, err = h.services.GetAllBooks(pagination)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllBooksResponse{
		Data:     books,
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	})

}

func (h *Handler) getBooksByFilter(c *gin.Context) {

	var pagination dtos.PaginationParams
	var err error
	pagination.Page, pagination.PageSize, err = dtos.ValidatePage(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var filter dtos.BookFilter
	if err := c.BindJSON(&filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = filter.Validate()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	books, err := h.services.GetBooksByFilter(filter, pagination)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Data:     books,
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	})

}

func (h *Handler) getBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	book, err := h.services.GetBookById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input dtos.BookUpdateInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = input.Validate()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateBook(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.DeleteBook(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
