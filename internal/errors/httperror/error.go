package httperror

import (
	"encoding/json"
	"fmt"
)

type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (h HTTPError) Error() string {
	if h.Cause == nil {
		return h.Detail
	}
	return h.Detail + " : " + h.Cause.Error()
}

func (h HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(h)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func (h *HTTPError) ResponseHeaders() (int, map[string]string) {
	return h.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func newHTTPError(err error, detail string, status int) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}
