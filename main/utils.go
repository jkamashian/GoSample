package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

// The response object for FooGet & FooPost
// foo data structure  2 fields. “name” field and  “id” field.
// Both string data types.
type FooData struct {
	Name string `json:"name"`
	Id   string `json:"id,omitempty"`
}

var Redis *redis.Client

func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	// Used for generic error response with {"Message": "your message here"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func ValidResponse(w http.ResponseWriter, uuid string, name string) {
	// used for valid responses for FooGet & FooPost
	// returns the  FooData 2 fields. A “name” field and an “id” field. Both are string data types.
	resp := FooData{
		Id:   uuid,
		Name: name,
	}
	jsonResp, respMarshalErr := json.Marshal(resp)
	if respMarshalErr != nil {
		panic(respMarshalErr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
