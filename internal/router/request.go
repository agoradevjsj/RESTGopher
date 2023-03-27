package router

import (
	"net/http"
	"strings"
)

type Request struct {
	path   string
	url    string
	params Params
	body   interface{}
}

type Params map[string]interface{}

type Item interface{}

func (r *RouterGopher) GetRequest(req *http.Request) *Request {
	request := &Request{
		path:   req.URL.Path,
		url:    req.URL.String(),
		body:   req.Body,
		params: make(map[string]interface{}),
	}
	request.extractParams(r.routes)
	return request
}

func (req *Request) extractParams(routes []RouteGopher) {
	for _, route := range routes {
		match, _ := ValidPathRouter(route.Path, req.path)
		if match {
			req.addParam(route.Path)
			break
		}
	}
}

func (req *Request) addParam(route string) {

	routeParts := strings.Split(route, "/")
	pathParts := strings.Split(req.path, "/")

	if len(routeParts) == len(pathParts) {
		for i, part := range routeParts {
			if strings.HasPrefix(part, ":") {
				paramName := strings.TrimSuffix(strings.TrimPrefix(part, ":"), "?")
				req.params[paramName] = pathParts[i]
			}
		}
	}
}

func (req *Request) GetPath() string {
	return req.path
}

func (req *Request) GetBody() interface{} {
	return req.body
}

func (req *Request) GetURL() string {
	return req.url
}

func (req *Request) GetParams() Params {
	return req.params
}

func (req *Request) GetJSON() map[string]interface{} {
	return req.params
}
