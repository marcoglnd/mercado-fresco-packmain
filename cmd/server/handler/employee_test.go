package controllers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/routes"
	"github.com/stretchr/testify/assert"
)

func getPathUrl(url string) string {
	PATH := "/api/v1"
	return fmt.Sprintf("%s%s", PATH, url)
}

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routerGroup := router.Group(getPathUrl(""))
	routes.AddRoutes(routerGroup)

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

func TestCreateOk(t *testing.T) {
	routes := createServer()
	req, res := createRequestTest(http.MethodPost, getPathUrl("/employees/"),
		` 
		{
			"card_number_id": "1234",
			"first_name": "Paloma",
			"last_name": "Ribeiro",
			"warehouse_id": 2
		}
		`,
	)
	defer req.Body.Close()
	routes.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}
