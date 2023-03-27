# RESTGopher

RESTGopher is a Golang library that makes creating and using REST APIs fun and accessible for Golang developers. With RESTGopher, developers can easily create and consume REST APIs, using a friendly and well-documented API.

RESTGopher is a simple and lightweight Go library for building RESTful APIs. It provides a router and middleware system, making it easy to define and manage routes, and to add functionality such as logging, authentication, and rate limiting to your API.

## Installation

To use RESTGopher, you must first install Go and set up your environment. Once that's done, you can install the library using the **`go get`** command:

```powershell
go get -u github.com/agoradevjsj/RESTGopher
```

## Usage

Here's an example of how to use RESTGopher to define and run a simple API:

```go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/agoradevjsj/RESTGopher"
	"github.com/agoradevjsj/RESTGopher/internal/middleware"
	"github.com/agoradevjsj/RESTGopher/internal/router"
)

func main() {
	// Create a new server
	app := restgopher.NewServerGopher(restgopher.CreateConfigApp(":4545", "", ""))

	// Define a route
	router_gopher := router.NewRouterGopher()
	router_gopher.GET("/hello/:wave", func(w http.ResponseWriter, r *http.Request) {
		// Get the content of Request, Response, and Params
		req := router_gopher.GetRequest(r)
		res := router_gopher.GetResponse(&w)
		params := req.GetParams()

		// Create a JSON format to display it in the response
		response := router.CreateJSON()

		response["ok"] = true
		response["http_code"] = 12000
		response["message"] = "This is my wave for ther world. " + params["wave"]
		// Set the message status and send the generated JSON format
		res.Status(http.StatusOK).SendJSON(response)
	})

	// Add middleware to the server
	app.Use(middleware.JSONMiddleware)
	app.Use(router_gopher.RoutersGopher)

	defer func() {
		r := recover()
		fmt.Println("⏺️ Recovered:", r)
	}()
	
	// Start the server
	errores := app.StartServerGopher()

	// Catch each error on the server
	if errores != nil {
		panic("Error on the server.")
	}
}
```

## API Reference

RESTGopher provides two main components: a router and a middleware system.

### Router

The router allows you to define and manage routes for your API. Here's an example of how to define a simple route using the router:

```go
// Define a new router
router_gopher := router.NewRouterGopher()

// Define a route
router_gopher.HandleFunc("GET", "/hello", func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello, HTTP!\n")
})

router_gopher.GET("/example", func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello, This a simple example!\n")
})

```

You can use any of the HTTP methods (GET, POST, PUT, DELETE, OPTIONS, PATCH) to define a route. The second parameter is the path for the route, and the third parameter is a function that will be called when the route is matched.

Middleware
RESTGopher provides a middleware system that allows you to add functionality to your API. Middleware are functions that take an **`http.HandlerFunc`** and return a new **`http.HandlerFunc`** that can modify the request or response. Here's an example of how to add JSON middleware to your server:

```go
// Add JSON middleware to the server
app.Use(middleware.JSONMiddleware)
```

This will add a middleware function that will automatically parse incoming JSON requests and set the **\`Content-Type\`** header to **\`application/json\`** for outgoing responses.

## License

RESTGopher is licensed under the MIT license. See the **[LICENSE](./LICENSE)** file for more details.