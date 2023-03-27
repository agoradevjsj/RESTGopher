package router

import (
	"net/http"
	"strings"
)

type RouteGopher struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type RouterGopher struct {
	routes []RouteGopher
}

func NewRouterGopher() *RouterGopher {
	return &RouterGopher{routes: make([]RouteGopher, 0)}
}

func (r *RouterGopher) HandleFunc(method string, path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{method, path, handler})
	return nil
}

func (r *RouterGopher) GET(path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{http.MethodGet, path, handler})
	return nil
}

func (r *RouterGopher) POST(path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{http.MethodPost, path, handler})
	return nil
}

func (r *RouterGopher) PUT(path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{http.MethodPut, path, handler})
	return nil
}

func (r *RouterGopher) DELETE(path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{http.MethodDelete, path, handler})
	return nil
}

func (r *RouterGopher) OPTIONS(path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{http.MethodOptions, path, handler})
	return nil
}

func (r *RouterGopher) PATCH(path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{http.MethodPatch, path, handler})
	return nil
}

func (r *RouterGopher) ANY(path string, handler http.HandlerFunc) error {
	r.routes = append(r.routes, RouteGopher{"ANY", path, handler})
	return nil
}

func ValidPathRouter(route string, path string) (matched bool, err error) {

	routeParts := strings.Split(route, "/")
	pathParts := strings.Split(path, "/")

	if len(routeParts) != len(pathParts) {
		return false, nil
	}

	for i := range routeParts {
		if routeParts[i] != pathParts[i] && !strings.HasPrefix(routeParts[i], ":") {
			return false, nil
		}
	}

	return true, nil
}

func (r *RouterGopher) RoutersGopher(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		for _, route := range r.routes {
			matched, _ := ValidPathRouter(route.Path, req.URL.Path)
			if route.Method == req.Method && matched {
				route.Handler(w, req)
				return
			} else if route.Method == "ANY" && matched {
				route.Handler(w, req)
				return
			}
		}

		http.Error(w, req.Method+" Not Found "+req.URL.Path, http.StatusNotFound)
	}
}
