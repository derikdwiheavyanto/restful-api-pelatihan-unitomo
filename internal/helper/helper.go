package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Meta struct {
}

func APIResponse(message string, code int, status string, data interface{}) Response {

	jsonResponse := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return jsonResponse
}

func FormatValidattionError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}

func IFErr(message string, code int, status string, data interface{}, c *gin.Context) {

	response := APIResponse(message, code, status, data)
	c.JSON(http.StatusBadRequest, response)

}

func BindJsonWhenSuccess(message string, status string, data interface{}, c *gin.Context) {
	response := APIResponse(message, http.StatusOK, status, data)
	c.JSON(http.StatusOK, response)
}
