package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

func GetCategories(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "GetCategories"})
	return http.StatusOK, result
}

func PostCategory(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "PostCategory"})
	return http.StatusOK, result
}

func GetCategory(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "GetCategory"})
	return http.StatusOK, result
}

func PutCategory(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "PutCategory"})
	return http.StatusOK, result
}

func DeleteCategory(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "DeleteCategory"})
	return http.StatusOK, result
}
