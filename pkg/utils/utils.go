package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//when we receive request in JSON format, we'll need to unmarshall it
//to make it understandable for our go program


//we're using io package to read the request body and store in type []byte

func ParseBody(r *http.Request, x interface{}){
	if body, err := io.ReadAll(r.Body); err==nil{
		if err := json.Unmarshal([]byte(body), x); err!=nil{
			fmt.Printf("%s", err)
			return;
		}
	}
}