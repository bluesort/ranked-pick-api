package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	api_errors "github.com/carterjackson/ranked-pick-api/internal/api/errors"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/go-chi/chi/v5"
)

type Route struct {
	Router chi.Router
	Method string
	Path   string
}

func Get(router chi.Router, path string) *Route {
	return &Route{
		Router: router,
		Method: "GET",
		Path:   path,
	}
}

func Post(router chi.Router, path string) *Route {
	return &Route{
		Router: router,
		Method: "POST",
		Path:   path,
	}
}

func Put(router chi.Router, path string) *Route {
	return &Route{
		Router: router,
		Method: "PUT",
		Path:   path,
	}
}

func Delete(router chi.Router, path string) *Route {
	return &Route{
		Router: router,
		Method: "DELETE",
		Path:   path,
	}
}

func (route *Route) Handler(handler interface{}, paramStruct ...interface{}) {
	routeHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var resp interface{}
		switch h := handler.(type) {
		case func(*common.Context) (interface{}, error):
			ctx, err := common.NewContext(r.Context())
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}

			resp, err = h(ctx)
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
		case func(*common.Context, interface{}) (interface{}, error):
			if len(paramStruct) == 0 {
				api_errors.WriteError(w, fmt.Errorf("missing paramStruct for path '%s'", route.Path))
				return
			}

			ctx, err := common.NewContext(r.Context())
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}

			params, err := extractParams(r, paramStruct[0])
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
			resp, err = h(ctx, params)
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
		case func(*common.Context, *db.Queries, interface{}) (interface{}, error):
			if len(paramStruct) == 0 {
				api_errors.WriteError(w, fmt.Errorf("missing paramStruct for path '%s'", route.Path))
				return
			}

			ctx, err := common.NewContext(r.Context())
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}

			params, err := extractParams(r, paramStruct[0])
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}

			tx, err := config.Config.Db.BeginTx(ctx, nil)
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
			defer tx.Rollback()
			txQueries := config.Config.Queries.WithTx(tx)

			resp, err = h(ctx, txQueries, params)
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}

			err = tx.Commit()
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
		case func(*common.Context, int64) (interface{}, error):
			ctx, err := common.NewContext(r.Context())
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}

			idStr := chi.URLParam(r, "id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				api_errors.WriteError(w, "Invalid id")
				return
			}

			resp, err = h(ctx, id)
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
		case func(*common.Context, *db.Queries, int64) error:
			ctx, err := common.NewContext(r.Context())
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}

			idStr := chi.URLParam(r, "id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				api_errors.WriteError(w, "Invalid id")
				return
			}

			tx, err := config.Config.Db.BeginTx(ctx, nil)
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
			defer tx.Rollback()
			txQueries := config.Config.Queries.WithTx(tx)

			err = h(ctx, txQueries, id)
			if err != nil {
				api_errors.WriteError(w, err)
				return
			}
		default:
			api_errors.WriteError(w, errors.New("unrecognized handler"))
			return
		}

		if resp != nil {
			json.NewEncoder(w).Encode(resp)
		}
		w.WriteHeader(http.StatusOK)
	}

	switch route.Method {
	case "GET":
		route.Router.Get(route.Path, routeHandler)
	case "POST":
		route.Router.Post(route.Path, routeHandler)
	case "PUT":
		route.Router.Put(route.Path, routeHandler)
	case "DELETE":
		route.Router.Delete(route.Path, routeHandler)
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
		return nil, api_errors.NewInputError("missing request body")
	} else if err != nil {
		return nil, err
	}

	return params, nil
}
