package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Data    string `json:",omitempty"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloworldhandler)

	log.Printf("server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloworldhandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("requestName: %s", request.Name)
	response := helloWorldResponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
