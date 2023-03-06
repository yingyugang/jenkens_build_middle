package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Empty struct {
}

func main() {
	http.HandleFunc("/build", BuildTrigger)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func BuildTrigger(w http.ResponseWriter, r *http.Request) {
	var empty Empty
	result, err := json.Marshal(empty)
	if err != nil {
		fmt.Printf("returnUser:get hero info error [%s]", err)
	}
	fmt.Println(string(result))
	w.Write(result)
}
