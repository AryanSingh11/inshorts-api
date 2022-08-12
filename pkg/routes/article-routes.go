package routes

import (
	"fmt"

	"github.com/AryanSingh11/inshorts-api/pkg/controllers"
	"github.com/gorilla/mux"
	//"github.com/AryanSingh11/inshorts-api/cmd/main"
)

var RegisterInshortsApiRoutes = func(router *mux.Router) {
	router.HandleFunc("/articles/", controllers.GetAllArticles).Methods("GET")
	router.HandleFunc("/articles/", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", controllers.GetArticleById).Methods("GET")
	router.HandleFunc("/articles/{id}", controllers.DeleteArticleById)
	//router.HandleFunc("/articles/search?q=<search text here>")

}
