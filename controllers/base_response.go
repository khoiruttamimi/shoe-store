package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data"`
	Pagination   *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	CurrentPage   int `json:"current_page"`
	LastPage      int `json:"last_page"`
	Count         int `json:"count"`
	RecordPerPage int `json:"record_per_page"`
}

func PaginationRes(page, count, limit int) Pagination {
	lastPage := count / limit
	if count%limit > 0 {
		lastPage = lastPage + 1
	}

	pagination := Pagination{
		CurrentPage:   page,
		LastPage:      lastPage,
		Count:         count,
		RecordPerPage: limit,
	}
	return pagination
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Data = data

	return c.JSON(http.StatusOK, response)
}
func NewSuccessResponseWithPagination(c echo.Context, data interface{}, pagination Pagination) error {
	response := BaseResponse{}
	response.Data = data
	response.Pagination = &pagination

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, err error) error {
	response := BaseResponse{}
	response.ErrorMessage = err.Error()
	if code, errCode := ErrorCode[err.Error()]; errCode {
		return c.JSON(code, response)
	}
	return c.JSON(500, response)
}
