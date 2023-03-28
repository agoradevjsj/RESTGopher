package router

import (
	"encoding/json"
	"io"
	"net/http"
)

// This a structure of Response
type Response struct {
	body   interface{}
	status int
	write  http.ResponseWriter
}

// Get the response RESTGopher.
func (r *RouterGopher) GetResponse(res *http.ResponseWriter) *Response {
	return &Response{
		status: 200,
		write:  *res,
	}
}

// Send data in text format
func (res *Response) SendText(text string) (int, error) {
	res.body = text
	res.write.WriteHeader(res.status)
	return io.WriteString(res.write, text)
}

// Change Status of your response
// and return the same response for you continue your sending data.
func (res *Response) Status(status int) *Response {
	res.status = status
	return res
}

// Send data in a Map on http(s) protcol.
func (res *Response) SendJSON(mapJSON map[string]interface{}) {
	res.body = mapJSON
	transformJSON, _ := json.Marshal(mapJSON)
	response := string(transformJSON)
	res.write.WriteHeader(res.status)
	io.WriteString(res.write, response)
}

// Send a simple text in format of Error
func (res *Response) SendError(text string) {
	res.body = text
	http.Error(res.write, text, res.status)
}

// Create a void map of a form faster than normally
func CreateJSON() map[string]interface{} {
	return make(map[string]interface{})
}

// Get status response
func (res *Response) GetStatus() int {
	return res.status
}

// Get body response
func (res *Response) GetBody() interface{} {
	return res.body
}
