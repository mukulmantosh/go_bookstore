package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	reqBody, _ := io.ReadAll(r.Body)
	r.Body.Close()
	err := json.Unmarshal(reqBody, x)
	if err != nil {
		return
	}
}
