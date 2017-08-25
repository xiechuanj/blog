package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

func GetTopics(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "GetTopics"})
	return http.StatusOK, result
}

func PostTopic(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "PostTopic"})
	return http.StatusOK, result
}

func GetTopic(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "GetTopic"})
	return http.StatusOK, result
}

func PutTopic(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "PutTopic"})
	return http.StatusOK, result
}

func DeleteTopic(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "DeleteTopic"})
	return http.StatusOK, result
}
