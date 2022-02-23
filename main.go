package main

import (
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	foo "github.com/jkamashian/GoSample/foo"
)

func main() {
	// simple webservice in go that will store “foo” records in memory(redis)
	redis := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err := redis.Ping().Result()
	if err != nil {
		// Failed to connect to redis
		panic(err)
	}

	//api routes
	router := mux.NewRouter()
	foo.InitializeFooEndpoints(router, redis)

	// The webservice runs on port 8080
	http.ListenAndServe(":8080", router)
}
