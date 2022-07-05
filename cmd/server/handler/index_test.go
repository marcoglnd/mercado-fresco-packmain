package controllers_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/routes"
)

func getPathUrl(url string) string {
	PATH := "/api/v1"
	return fmt.Sprintf("%s%s", PATH, url)
}

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routerGroup := router.Group(getPathUrl(""))
	routes.AddRoutes(routerGroup, &sql.DB{})

	return router
}

func createRequestTest(
	method string,
	url string,
	body string,
) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}
