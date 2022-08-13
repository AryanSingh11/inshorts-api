package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/AryanSingh11/inshorts-api/pkg/models"
	"github.com/gorilla/mux"
)



func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fileList, err := os.ReadDir("../../data")
	if err != nil {
		fmt.Println(err)
	}

	//Now we will range over each .json file in fileList
	//and read data from all of them and write
	//them on our response
	for _, file := range fileList {

		//fmt.Println(file.Name())

		fileString := "../../data/" + file.Name()

		fileData, err := os.ReadFile(fileString)
		if err != nil {
			fmt.Println(err)
		}

		w.Write(fileData)
		w.Write([]byte("\n"))

	}

}


func CreateArticle(w http.ResponseWriter, r *http.Request) {
	//io.ReadAll gives output in byte format
	//json.Unmarshal accepts first arg in []byte format
	jsonBodyInByte, _ := io.ReadAll(r.Body)
	var article models.Article
	err := json.Unmarshal(jsonBodyInByte, &article)
	if err != nil {
		log.Fatal(err)
	}
	//getting id from req body
	extractedId := article.Id
	fileString := "../../data/" + strconv.Itoa(int(extractedId)) + ".json"
	os.WriteFile(fileString, jsonBodyInByte, 0666)

	//sending response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	respString := "post request received for article id: " + strconv.Itoa(int(extractedId))
	w.Write([]byte(respString))

}



func GetArticleById(w http.ResponseWriter, r *http.Request) {

	//first step is to extract the id
	vars := mux.Vars(r)
	articleId := vars["id"]
	Id, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		log.Fatal("error converting the Id to string")
	}
	//fmt.Println(Id)


	//second step is to get the corresponding file
	//when we get the file, read the file and write
	//its content as response
	fileName := strconv.Itoa(int(Id))+".json"
	fileList, err := os.ReadDir("../../data")
	if err != nil {
		fmt.Println(err)
	}

	//we'll match the id with the list of files
	//we have. Also i is here to help us track if 
	//some data was sent or not
	var i int = 0
	for _, file := range fileList {

		if(file.Name()==fileName){
			fileString := "../../data/" + fileName
			fileData, err := os.ReadFile(fileString)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(fileData)
			w.Write([]byte("\n"))
			i++
			return 
		}

	}
	//if not id matches return a message
	if(i==0){
		w.Write([]byte("sorry no article with this Id was found😊"))
	}

}


func SearchForArticle(w http.ResponseWriter, r *http.Request){

	//first we will extract the search text
	//below key searches for first key q in url string
	searchText := r.URL.Query().Get("q")

	//we'll now iterate over each file in data dir
	//then run our search in each file
	fileList, err := os.ReadDir("../../data")
	if err != nil {
		fmt.Println(err)
	}

	var i int = 0
	for _, file := range fileList{

		fileString := "../../data/"+file.Name()
		fmt.Println(fileString)
		fileData, err := os.ReadFile(fileString)
		if err != nil {
			fmt.Println(err)
		}

		//if strings in fileData contains our searchText
		//send that fileData as response
		if( strings.Contains( string(fileData) , searchText)){
			w.Write(fileData)
			w.Write([]byte("/n"))
			i++
		}

	}
	//msg for no matching string
	if(i==0){
		w.Write([]byte("no matching article was found 😓"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}


func DeleteArticleById(w http.ResponseWriter, r *http.Request) {

	//extracting Id from url string
	vars := mux.Vars(r)
	articleId := vars["id"]
	Id, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		log.Fatal("error converting the Id to string")
	}

	//second step is to get the corresponding file
	//when we get the file, read the file and write
	//its content as response
	fileName := strconv.Itoa(int(Id))+".json"
	fileList, err := os.ReadDir("../../data")
	if err != nil {
		fmt.Println(err)
	}

	//we'll match the id with the list of files
	//we have. Also i is here to help us track if 
	//some file was deleted or not
	var i int = 0
	for _, file := range fileList {

		if(file.Name()==fileName){
			fileString := "../../data/" + fileName
			err := os.Remove(fileString)
			if err != nil {
				fmt.Println(err)
			}
			deleteMessage := "article delted with id : "+ strconv.Itoa(int(Id))
			w.Write([]byte(deleteMessage))
			i++
			return 
		}

	}
	//if no id matches return a message
	if(i==0){
		w.Write([]byte("sorry no article with this Id was found to delete😊"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
