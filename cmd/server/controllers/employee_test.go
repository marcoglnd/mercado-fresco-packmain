package controllers_test

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	return r
}

func Test_CreateEmployee_Ok(t *testing.T) {
	print("Ok")
}
