package main

import (
	"RetrieveData/Tables"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	r.Use(middlewareLogger)
	r.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request){

		res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Println(res.Body)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		defer res.Body.Close()

		var uResp []Tables.User

		if err := json.Unmarshal(body, &uResp); err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		bite, err := json.Marshal(uResp)
		if err!=nil{
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Fprintf(writer, string(bite))
		// fmt.Printf("%+v\n", uResp)
	})
	fmt.Printf("Alive and running at localhost:4097")
	http.ListenAndServe(":4097", r)
	
}

func middlewareLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("accessing %v\n", request.URL)
		handler.ServeHTTP(writer, request)
	})
}