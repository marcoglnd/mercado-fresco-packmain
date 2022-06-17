package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/cmd/server/controllers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := buyers.NewRepository()
	service := buyers.NewService(repo)

	buyerController := controllers.NewBuyer(service)

	router := gin.Default()

	pr := router.Group("/buyers")
	{
		pr.GET("/", buyerController.GetAll())
		pr.POST("/", buyerController.Create())
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

func Test_CreateBuyer_OK(t *testing.T) {
	// crie o Servidor e defina as Rotas
	r := createServer()
	// crie Request do tipo POST e Response para obter o resultado
	req, rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"card_number_id": "402323", "first_name": "Jhon", "last_name": "Doe"
	}`)

	// diga ao servidor que ele pode atender a solicitação
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func Test_CreateBuyer_fail(t *testing.T) {
	// crie o Servidor e defina as Rotas
	r := createServer()
	// crie Request do tipo POST e Response para obter o resultado
	req, rr := createRequestTest(http.MethodPost, "/buyers/", `{
		"first_name": "Jhon", "last_name": "Doe"
	}`)

	// diga ao servidor que ele pode atender a solicitação
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func Test_GetBuyers_OK(t *testing.T) {
	// criar um servidor e define suas rotas
	r := createServer()
	// criar uma Request do tipo GET e Response para obter o resultado
	req, rr := createRequestTest(http.MethodGet, "/buyers/", "")

	defer req.Body.Close()

	// diz ao servidor que ele pode atender a solicitação
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	objRes := struct {
		Code int
		Data []buyers.Buyer
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}
