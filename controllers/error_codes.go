package controller

import "net/http"

var ErrorCode = map[string]int{
	"data not found":                   http.StatusBadRequest,
	"duplicate data":                   http.StatusBadRequest,
	"id not found":                     http.StatusBadRequest,
	"(ProductCode) not found or empty": http.StatusBadRequest,
	"(ProductID) not found or empty":   http.StatusBadRequest,
	"invalid parameter id":             http.StatusBadRequest,
	"invalid admin key":                http.StatusBadRequest,
	"record not found":                 http.StatusBadRequest,
	"(Phone) or (Password) empty":      http.StatusBadRequest,
	"invalid credential":               http.StatusUnauthorized,
	"something gone wrong":             http.StatusInternalServerError,
	"forbidden roles":                  http.StatusForbidden,
}
