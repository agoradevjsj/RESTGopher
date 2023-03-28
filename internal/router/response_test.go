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

func TestGetResponse(t *testing.T) {
	server := restgopher.NewServerGopher(restgopher.CreateConfigApp(":4545", "", "Hello World!!"))
	router_gopher := router.NewRouterGopher()
	assert.NotNil(t, router_gopher, "Routes should have being Not Nil")
	router_gopher.GET("/hello/:wave", func(w http.ResponseWriter, r *http.Request) {
		req := router_gopher.GetRequest(r)
		assert.NotNil(t, req, "Request should have being Not Nil")

		res := router_gopher.GetResponse(&w)
		assert.NotNil(t, res, "Response should have being Not Nil")

		params := req.GetParams()
		assert.NotNil(t, params, "Params should have being Not Nil")

		urlTest := req.GetURL()
		assert.NotSame(t, "", urlTest, "The URL shouldn't have being void but its value is ", urlTest)

		pathTest := req.GetPath()
		assert.NotSame(t, "", pathTest, "The Path shouldn't have being void but its value is ", pathTest)

		jsonTest := req.GetJSON()
		assert.NotNil(t, jsonTest, "Params json should have being Not Nil")

		assert.NotEmpty(t, params["wave"], "Param 'wave' should have being Not Nil")
		response := router.CreateJSON()
		assert.Equal(t, make(map[string]interface{}), response, "The JSON should have being a new json format")

		response["ok"] = true
		response["message"] = "This is my wave... " + fmt.Sprintf("%v", params["wave"])

		res.Status(http.StatusOK).SendJSON(response)
		assert.Equal(t, http.StatusOK, res.GetStatus(), "The status should have being 200 OK but it's a status", res.Status(http.StatusOK))
		assert.Equal(t, response, res.GetBody(), "The status should have being 200 OK but it's a status", res.Status(http.StatusOK))
		textTest := "This a test for simple send text"

		res.Status(http.StatusOK).SendText(textTest)
		assert.Equal(t, http.StatusOK, res.GetStatus(), "The status should have being 200 OK but it's a status", res.Status(http.StatusOK))
		assert.Equal(t, textTest, res.GetBody(), "The status should have being 200 OK but it's a status", res.Status(http.StatusOK))
		textTest = "This a test for simple send error"

		res.Status(http.StatusBadRequest).SendError(textTest)
		assert.Equal(t, http.StatusBadRequest, res.GetStatus(), "The status should have being 400 Bad Request but it's a status", res.Status(http.StatusOK))
		assert.Equal(t, textTest, res.GetBody(), "The status should have being 200 OK but it's a status", res.Status(http.StatusOK))
	})
	server.Use(middleware.JSONMiddleware, router_gopher.RoutersGopher)

	go server.StartServerGopher()
	time.Sleep(1000 * time.Millisecond)
	res, err := http.Get("http://localhost:4545/hello/hello-world")
	assert.Nil(t, err, "The error should have being nil but there is a error:", err)
	assert.NotNil(t, res, "The response should have being Not Nil but is Nil", res)
}
