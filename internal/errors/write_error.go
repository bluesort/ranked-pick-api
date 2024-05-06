package errors

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func WriteError(w http.ResponseWriter, err interface{}) {
	switch errVal := err.(type) {
	case *InputError:
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("\"" + errVal.Message + "\""))
	case error:
		log.Print(errVal)
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
