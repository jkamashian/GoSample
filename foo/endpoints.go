package foo

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func PostFoo(w http.ResponseWriter, r *http.Request) {
	// support POST to endpoint (‘/foo’)
	// accepts a json foo object and responds with a 200 response code
	// $ curl -i -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' 'http://localhost:8080/foo' -d '{"name": "Jack"}'
	var fn FooData
	err := json.NewDecoder(r.Body).Decode(&fn)
	if err != nil {
		ErrorResponse(w, "Invalid Request", 400)
		return
	}
	uuidString := uuid.NewString()
	reddisErr := Redis.Set(uuidString, fn.Name, 0).Err()
	if reddisErr != nil {
		ErrorResponse(w, "Server Error", 500)
		return
	}
	ValidResponse(w, uuidString, fn.Name)
}

func GetFoo(w http.ResponseWriter, r *http.Request) {
	// GET endpoint (‘foo/{id}’) that responds with a 200 response code if the record is found
	// or a 404 response code if not found.
	//$ curl -i -X GET -H 'Accept: application/json' 'http://localhost:8080/foo/26baf48a-db0f-4884-9b89-820ce7596a6e'
	urlParams := mux.Vars(r)
	val, err := Redis.Get(urlParams["id"]).Result()
	if err != nil {
		w.WriteHeader(404)
		return
	}
	ValidResponse(w, urlParams["id"], val)
}

func DeleteFoo(w http.ResponseWriter, r *http.Request) {
	// DELETE endpoint (‘foo/{id}’) that responds with a 204 response code if the record is found
	// or a 404 response code if not found.
	urlParams := mux.Vars(r)
	val, err := Redis.Del(urlParams["id"]).Result()
	if err != nil {
		ErrorResponse(w, "Server Error", 500)
		return
	}
	if val == 1 {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(404)
	}
}
