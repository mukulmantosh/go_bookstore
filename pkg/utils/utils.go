package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
		return
	}

	//if body, err :=  err == nil {
	//	if err := json.Unmarshal([]byte(body), x); err != nil {
	//		return
	//	}
	//}
}
