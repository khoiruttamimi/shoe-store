package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
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
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}
func NewSuccessResponseWithPagination(c echo.Context, data interface{}, pagination Pagination) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data
	response.Pagination = &pagination

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
