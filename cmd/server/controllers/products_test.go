package controllers_test

import (
	"bytes"
	"encoding/json"
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

func TestCreateProductUnprocessable(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/products/", `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"freezing_rate": 2,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	second_req, second_rr := createRequestTest(http.MethodPost, "/products/", `{
		"description": "Yogurt",
		"expiration_rate": 1,
		"height": 6.4,
		"length": 4.5,
		"netweight": 3.4,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 1.3,
		"width": 1.2,
		"product_type_id": 2,
		"seller_id": 2
		}`)

	third_req, third_rr := createRequestTest(http.MethodPost, "/products/", `{
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
	r.ServeHTTP(second_rr, second_req)
	r.ServeHTTP(third_rr, third_req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, second_rr.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, third_rr.Code)
}

func TestCreateProductConflict(t *testing.T) {
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

	second_req, second_rr := createRequestTest(http.MethodPost, "/products/", `{
		"description": "Queijo",
		"expiration_rate": 2,
		"freezing_rate": 3,
		"height": 8.6,
		"length": 2.4,
		"netweight": 5.7,
		"product_code": "PROD01",
		"recommended_freezing_temperature": 4.5,
		"width": 2.5,
		"product_type_id": 54,
		"seller_id": 1
		}`)

	r.ServeHTTP(rr, req)
	r.ServeHTTP(second_rr, second_req)

	assert.Equal(t, http.StatusConflict, second_rr.Code)
}

func TestGetAllOK(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	defer req.Body.Close()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []products.Product
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}
