package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type fruit struct {
	ID    int    `json:id`
	Name  string `json:name`
	Color string `json:color`
}

var fruits = []fruit{
	{ID: 1, Name: "apple", Color: "red"},
	{ID: 2, Name: "banana", Color: "yellow"},
	{ID: 3, Name: "orange", Color: "orange"},
}

func getFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fruits)
}

func main() {
	http.HandleFunc("/", getFruits)
	fmt.Println("start server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
