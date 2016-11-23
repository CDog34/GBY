package Model

import "net/http"

type ErrorDetail struct {
	HttpCode  int
	ErrorCode int
	Message   string
}

var Errors = map[string]ErrorDetail{
	"regular.notFound":{http.StatusNotFound, 10000, "ResourceNotFound"},
	"not found":{http.StatusNotFound, 10000, "ResourceNotFound"},

	"paramErr.validFail":{http.StatusBadRequest, 20000, "ParamsValidationFail"},


}