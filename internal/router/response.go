package router

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	body   interface{}
	status int
	write  http.ResponseWriter
}

func (r *RouterGopher) GetResponse(res *http.ResponseWriter) *Response {
	return &Response{
		status: 200,
		write:  *res,
	}
}

func (res *Response) SendText(text string) (int, error) {
	res.write.WriteHeader(res.status)
	return io.WriteString(res.write, text)
}

func (res *Response) Status(status int) *Response {
	res.status = status
	return res
}

func (res *Response) SendJSON(mapJSON map[string]interface{}) {
	res.body = mapJSON
	transformJSON, _ := json.Marshal(mapJSON)
	response := string(transformJSON)
	res.write.WriteHeader(res.status)
	io.WriteString(res.write, response)
}

func (res *Response) SendError(text string) {
	http.Error(res.write, text, res.status)
}

func CreateJSON() map[string]interface{} {
	return make(map[string]interface{})
}
