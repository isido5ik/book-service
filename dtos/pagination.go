package dtos

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationParams struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func ValidatePage(c *gin.Context) (int, int, error) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "3")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return -1, -1, errors.New("invalid value of query parameter 'page' (it must by integer)")
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return -1, -1, errors.New("invalid value of query parameter 'pageSize' (it must by integer)")
	}

	if pageSize <= 0 || page <= 0 {
		return -1, -1, errors.New("invalid value og query parameter 'page' or 'pageSize', negative number or zero")
	}
	return page, pageSize, nil
}
