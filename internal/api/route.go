package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	router *chi.Mux
	method string
	path   string
}

func Get(router *chi.Mux, path string) *Route {
	return &Route{
		router: router,
		method: "GET",
		path:   path,
	}
}

func Post(router *chi.Mux, path string) *Route {
	return &Route{
		router: router,
		method: "POST",
		path:   path,
	}
}

func Put(router *chi.Mux, path string) *Route {
	return &Route{
		router: router,
		method: "PUT",
		path:   path,
	}
}

func (r *Route) Handler(handler interface{}, paramStruct ...interface{}) {
	routeHandler := func(w http.ResponseWriter, req *http.Request) {
		var resp interface{}
		var err error
		switch h := handler.(type) {
		case func(interface{}) (interface{}, error):
			resp, err = h(nil)
		case func(interface{}, interface{}) (interface{}, error):
			if len(paramStruct) == 0 {
				panic(fmt.Sprintf("Missing paramStruct for path '%s'", r.path))
			}

			param := extractParams(req, paramStruct[0])
			resp, err = h(nil, param)

		case func(interface{}, int64) (interface{}, error):
			// TODO
		default:
			panic("Unrecognized handler")
		}
		if err != nil {
			WriteError(w, err)
			return
		}

		json.NewEncoder(w).Encode(resp)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}

	switch r.method {
	case "GET":
		r.router.Get(r.path, routeHandler)
	case "POST":
		r.router.Post(r.path, routeHandler)
	case "PUT":
		r.router.Put(r.path, routeHandler)
	}
}

func extractParams(r *http.Request, paramStruct interface{}) interface{} {
	if r.Method == "GET" {
		paramType := reflect.ValueOf(paramStruct).Type()
		params := reflect.New(paramType).Interface()

		// encode and decode URL query params into the param struct
		urlParams := r.URL.Query()
		queryParams := make(map[string]interface{}, len(urlParams))
		for k, v := range urlParams {
			if len(v) == 1 {
				queryParams[k] = v[0]
			} else {
				queryParams[k] = v
			}
		}
		encodedQueryParams, err := json.Marshal(queryParams)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(encodedQueryParams, params)
		if err != nil {
			panic(err)
		}
		return params
	}

	// TODO: post params
	return nil
}
