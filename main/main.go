package main

import (
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func main() {
	// simple webservice in go that will store “foo” records in memory(redis)
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		// Failed to connect to redis
		panic(err)
	}

	//api routes
	router := mux.NewRouter()
	router.HandleFunc("/foo", PostFoo).Methods("POST")
	router.HandleFunc("/foo/{id}", GetFoo).Methods("GET")
	router.HandleFunc("/foo/{id}", DeleteFoo).Methods("DELETE")

	// The webservice runs on port 8080
	http.ListenAndServe(":8080", router)
}
