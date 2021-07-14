package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpResult struct {
	Data  interface{}
	Error string
}

func Error(w http.ResponseWriter, code int, errHTML error) {
	response(w, code, nil, errHTML)
	return
}

func response(w http.ResponseWriter, code int, data interface{}, errHTML error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(httpResult{
		Data:  data,
		Error: errHTML.Error(),
	})
	if err != nil {
		_, _ = fmt.Fprintf(w, "parse error - %s", err.Error())
	}
	return
}

func Json(w http.ResponseWriter, code int, data interface{}) {
	response(w, code, data, nil)
	return
}
