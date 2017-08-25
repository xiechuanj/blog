package handler

import (
	"encoding/json"
	"gopkg.in/macaron.v1"
	"net/http"
)

// IndexV1Handler is
func IndexV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "Blog Backend REST API Service"})
	return http.StatusOK, result
}
