package main

import (
	"log"
	"net/http"

	"github.com/ragul28/geo-tracking-redis/config"
	"github.com/ragul28/geo-tracking-redis/handler"
)

func main() {
	env := config.GetEnv()

	r := handler.InitRouter()

	log.Printf("Server started on Port: %v", env.Port)
	log.Fatal(http.ListenAndServe(":"+env.Port, r))
}
