package che

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpError struct {
	Date  interface{}
	Error string
}

func Error(w http.ResponseWriter, code int, errStr string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(httpError{
		Error: errStr,
	})
	if err != nil {
		_, _ = fmt.Fprintf(w, "parse error - %s", err.Error())
	}
	return
}
