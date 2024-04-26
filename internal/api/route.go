package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/db"
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
		ctx := common.NewContext()
		w.Header().Set("Content-Type", "application/json")

		var err error
		var resp interface{}
		switch h := handler.(type) {
		case func(*common.Context) (interface{}, error):
			resp, err = h(ctx)
			if err != nil {
				WriteError(w, err)
				return
			}
		case func(*common.Context, interface{}) (interface{}, error):
			if len(paramStruct) == 0 {
				WriteError(w, errors.New(fmt.Sprintf("Missing paramStruct for path '%s'", r.path)))
				return
			}

			params, err := extractParams(req, paramStruct[0])
			if err != nil {
				WriteError(w, err)
				return
			}
			resp, err = h(ctx, params)
			if err != nil {
				WriteError(w, err)
				return
			}
		case func(*common.Context, *db.Queries, interface{}) (interface{}, error):
			if len(paramStruct) == 0 {
				WriteError(w, errors.New(fmt.Sprintf("Missing paramStruct for path '%s'", r.path)))
				return
			}

			params, err := extractParams(req, paramStruct[0])
			if err != nil {
				WriteError(w, err)
				return
			}

			tx, err := config.Config.Db.BeginTx(ctx, nil)
			if err != nil {
				WriteError(w, err)
				return
			}
			defer tx.Rollback()
			txQueries := config.Config.Queries.WithTx(tx)

			resp, err = h(ctx, txQueries, params)
			if err != nil {
				WriteError(w, err)
				return
			}

			err = tx.Commit()
			if err != nil {
				WriteError(w, err)
				return
			}
		case func(*common.Context, int64) (interface{}, error):
			idStr := chi.URLParam(req, "id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				WriteError(w, "Invalid id")
				return
			}

			resp, err = h(ctx, id)
			if err != nil {
				WriteError(w, err)
				return
			}
		case func(*common.Context, *db.Queries, int64) error:
			idStr := chi.URLParam(req, "id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				WriteError(w, "Invalid id")
				return
			}

			tx, err := config.Config.Db.BeginTx(ctx, nil)
			if err != nil {
				WriteError(w, err)
				return
			}
			defer tx.Rollback()
			txQueries := config.Config.Queries.WithTx(tx)

			err = h(ctx, txQueries, id)
			if err != nil {
				WriteError(w, err)
				return
			}
		default:
			WriteError(w, errors.New("Unrecognized handler"))
			return
		}

		if resp != nil {
			json.NewEncoder(w).Encode(resp)
		}
		w.WriteHeader(http.StatusOK)
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

func extractParams(r *http.Request, paramStruct interface{}) (interface{}, error) {
	paramVal := reflect.ValueOf(paramStruct)
	if paramVal.Kind() == reflect.Ptr {
		paramVal = paramVal.Elem()
	}
	paramType := paramVal.Type()
	params := reflect.New(paramType).Interface()

	if r.Method == "GET" {
		// marshal and unmarshal URL query params into the param struct
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
			return nil, err
		}
		err = json.Unmarshal(encodedQueryParams, params)
		if err != nil {
			return nil, err
		}
		return params, nil
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err == io.EOF {
		return nil, NewInputError("Missing request body")
	} else if err != nil {
		return nil, err
	}

	return params, nil
}
