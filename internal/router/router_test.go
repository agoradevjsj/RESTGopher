package router_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	restgopher "github.com/agoradevjsj/RESTGopher"
	"github.com/agoradevjsj/RESTGopher/internal/middleware"
	"github.com/agoradevjsj/RESTGopher/internal/router"
	"github.com/stretchr/testify/assert"
)

type ResponseURL struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func TestValidHttpMethods(t *testing.T) {
	server := restgopher.NewServerGopher(restgopher.CreateConfigApp(":4545", "localhost", "Hello World!!"))
	router_gopher := router.NewRouterGopher()
	assert.NotNil(t, router_gopher, "Routes should have being Nil")
	httpGET := router_gopher.GET("/hello/:wave", func(w http.ResponseWriter, r *http.Request) {
		req := router_gopher.GetRequest(r)
		assert.NotNil(t, req, "Request should have being Not Nil")
		res := router_gopher.GetResponse(&w)
		assert.NotNil(t, res, "Response should have being Not Nil")
		params := req.GetParams()
		assert.NotNil(t, params, "Params should have being Not Nil")
		assert.NotEmpty(t, params["wave"], "Param 'wave' should have being Not Nil")

		response := router.CreateJSON()
		assert.Equal(t, make(map[string]interface{}), response, "The JSON should have being a new json format")

		response["ok"] = true
		response["message"] = "This is my wave... " + fmt.Sprintf("%v", params["wave"])

		res.Status(http.StatusOK).SendJSON(response)
	})
	assert.Nil(t, httpGET, "Http GET should have being Nil")

	httpPOST := router_gopher.POST("/hello/:wave", func(w http.ResponseWriter, r *http.Request) {
		req := router_gopher.GetRequest(r)
		assert.NotNil(t, req, "Request should have being Not Nil")
		res := router_gopher.GetResponse(&w)
		assert.NotNil(t, res, "Response should have being Not Nil")
		params := req.GetParams()
		assert.NotNil(t, params, "Params should have being Not Nil")
		assert.NotEmpty(t, params["wave"], "Param 'wave' should have being Not Nil")

		response := router.CreateJSON()
		assert.Equal(t, make(map[string]interface{}), response, "The JSON should have being a new json format")

		response["ok"] = true
		response["message"] = "This is my wave... " + fmt.Sprintf("%v", params["wave"])

		res.Status(http.StatusOK).SendJSON(response)
	})
	assert.Nil(t, httpPOST, "Http POST should have being Nil")

	httpDELETE := router_gopher.DELETE("/hello/:wave", func(w http.ResponseWriter, r *http.Request) {
		req := router_gopher.GetRequest(r)
		assert.NotNil(t, req, "Request should have being Not Nil")
		res := router_gopher.GetResponse(&w)
		assert.NotNil(t, res, "Response should have being Not Nil")
		params := req.GetParams()
		assert.NotNil(t, params, "Params should have being Not Nil")
		assert.NotEmpty(t, params["wave"], "Param 'wave' should have being Not Nil")

		response := router.CreateJSON()
		assert.Equal(t, make(map[string]interface{}), response, "The JSON should have being a new json format")

		response["ok"] = true
		response["message"] = "This is my wave... " + fmt.Sprintf("%v", params["wave"])

		res.Status(http.StatusOK).SendJSON(response)
	})
	assert.Nil(t, httpDELETE, "Http DELETE should have being Nil")

	httpPUT := router_gopher.PUT("/hello/:wave", func(w http.ResponseWriter, r *http.Request) {
		req := router_gopher.GetRequest(r)
		assert.NotNil(t, req, "Request should have being Not Nil")
		res := router_gopher.GetResponse(&w)
		assert.NotNil(t, res, "Response should have being Not Nil")
		params := req.GetParams()
		assert.NotNil(t, params, "Params should have being Not Nil")
		assert.NotEmpty(t, params["wave"], "Param 'wave' should have being Not Nil")

		response := router.CreateJSON()
		assert.Equal(t, make(map[string]interface{}), response, "The JSON should have being a new json format")

		response["ok"] = true
		response["message"] = "This is my wave... " + fmt.Sprintf("%v", params["wave"])

		res.Status(http.StatusOK).SendJSON(response)
	})
	assert.Nil(t, httpPUT, "Http PUT should have being Nil")

	server.Use(middleware.JSONMiddleware, router_gopher.RoutersGopher)

	go server.StartServerGopher()
	time.Sleep(1000 * time.Millisecond)
	res, err := http.Get("http://localhost:4545/hello/hello-world")
	assert.Nil(t, err, "The error GET should have being nil but there is a error:", err)
	assert.NotNil(t, res, "The response GET should have being Not Nil but is Nil", res)

	res, err = http.Post("http://localhost:4545/hello/hello-world", "", nil)
	assert.Nil(t, err, "The error GET should have being nil but there is a error:", err)
	assert.NotNil(t, res, "The response GET should have being Not Nil but is Nil", res)
}
