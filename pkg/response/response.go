package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Status:  10000,
		Message: msg,
		Data:    data,
	})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Status:  400,
		Message: msg,
		Data:    nil,
	})
}
