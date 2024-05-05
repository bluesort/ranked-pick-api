package api

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/carterjackson/ranked-pick-api/internal/api/errors"
)

func WriteError(w http.ResponseWriter, err interface{}) {
	switch errVal := err.(type) {
	case *errors.InputError:
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("\"" + errVal.Message + "\""))
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("\"error\""))
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
