package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
) 

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler
func handler(w http.ResponseWriter, r *http.Request){
	var j bytes.Buffer
	b, _ := ioutil.ReadAll(r.Body)
	json.Indent(&j, b, "", "  ")
	fmt.Fprintf(w, j.String())
}

// curl --header "Content-Type: application/json" --request POST --data '{"username":"xyz","password":"xyz"}' http://localhost:3000/api/login