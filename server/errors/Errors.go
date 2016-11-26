package errorHandler

import (
	"net/http"
	"strings"
)

type ErrorDetail struct {
	HttpCode  int    `json:"-"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

var appErrors = map[string]ErrorDetail{
	"regular.notFound": {http.StatusNotFound, 10000, "ResourceNotFound"},
	"not found":        {http.StatusNotFound, 10000, "ResourceNotFound"},

	"paramErr.validFail":          {http.StatusInternalServerError, 20000, "paramErr.validFail"},
	"paramErr.validFail.required": {http.StatusBadRequest, 20001, "ParamsValidationFail"},
	"paramErr.validFail.type":     {http.StatusBadRequest, 20002, "ParamsValidationFail"},
	"paramErr.inValidObjectId":    {http.StatusNotFound, 20003, "inValidObjectId"},

	"auth.needSession": {http.StatusUnauthorized, 30001, "needSession"},
}

func GetErrorDetail(err error) ErrorDetail {
	errArr := strings.Split(err.Error(), "/*/")
	errorKey := errArr[0]
	result := appErrors[errorKey]
	if result.HttpCode == 0 {
		result = ErrorDetail{http.StatusInternalServerError, 66666, "InternalServerError"}
	}
	if len(errArr) > 1 && errArr[1] != "" {
		result.Message = errArr[1]
	}
	return result
}
