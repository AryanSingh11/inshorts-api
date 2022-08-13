package main

import (
	"fmt"
	"github.com/AryanSingh11/inshorts-api/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterInshortsApiRoutes(router)
	http.Handle("/", router)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal(err)
	}

}
