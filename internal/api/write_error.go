package api

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

func WriteError(w http.ResponseWriter, err interface{}) {
	switch errVal := err.(type) {
	case *errors.InputError:
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(errorResp(errVal.Message))
	case *errors.NotFoundError:
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorResp(errVal.Error()))
	case *errors.AuthError:
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(errorResp(errVal.Error()))
	case error:
		log.Print(errVal)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResp("something went wrong"))
	default:
		var typeName string
		if t := reflect.TypeOf(err); t.Kind() == reflect.Ptr {
			typeName = t.Elem().Name()
		} else {
			typeName = t.Name()
		}
		panic(fmt.Sprintf("Invalid error of type '%s' returned", typeName))
	}
}

func errorResp(message string) []byte {
	return []byte("{\"error\":\"" + message + "\"}")
}
