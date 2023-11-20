package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpResponse struct {
	Status		int						`json:"status"`
	Message		string					`json:"message"`
	Data		map[string]string		`json:"data"`
}

func SendJSONResponse(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	fmt.Println("Sending Payload", payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}