package main

import (
	"fmt"
	"github.com/AryanSingh11/inshorts-api/pkg/routes"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	http.Handle("/", router)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
