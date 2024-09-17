package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielecr/nsa-golang-rest-mysql/model"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := model.TodosByUserId(1)
	fmt.Println(todos)
	jd, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s\n", jd)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	http.ListenAndServe(":8083", router)
	fmt.Println("Hello, World!")
}
