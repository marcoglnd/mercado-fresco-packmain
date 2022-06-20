package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/products"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := controllers.NewProduct(service)

	router := gin.Default()

	pr := router.Group("/products")
	{
		pr.GET("/", controller.GetAll())
		pr.GET("/:id", controller.GetById())
		pr.POST("/", controller.CreateNewProduct())
		pr.PATCH("/:id", controller.Update())
		pr.DELETE("/:id", controller.Delete())
	}

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

func TestCreateProductOK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/products/", `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
