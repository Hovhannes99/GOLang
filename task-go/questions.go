package main

import (
    "encoding/json"
    "net/http"
)
import "fmt"

type Answer struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Correct bool   `json:"correct"`
}


type Question struct {
	ID      int      `json:"id"`
	Content string   `json:"content"`
	Answers []Answer `json:"answers"`
}

func main() {

   questions := []Question{
   		{
   			1,
   			"What is the capital of France?",
   			[]Answer{
   				{1, "Berlin", false},
   				{2, "Madrid", false},
   				{3, "Rome", false},
   				{4, "Paris", true},
   			},
   		},
   		{
   			2,
   			"What is 2 + 2?",
   			[]Answer{
   				{1, "3", false},
   				{2, "4", true},
   				{3, "5", false},
   				{4, "6", false},
   			},
   		},
   		{
   			3,
   			"What is the largest planet in our solar system?",
   			[]Answer{
   				{1, "Earth", false},
   				{2, "Mars", false},
   				{3, "Jupiter", true},
   				{4, "Venus", false},
   			},
   		},
   		{
   			4,
   			"Which programming language is known for its simplicity and readability?",
   			[]Answer{
   				{1, "Java", false},
   				{2, "Python", true},
   				{3, "C++", false},
   				{4, "JavaScript", false},
   			},
   		},
   		{
   			5,
   			"What gas do plants absorb from the atmosphere?",
   			[]Answer{
   				{1, "Oxygen", false},
   				{2, "Carbon dioxide", true},
   				{3, "Nitrogen", false},
   				{4, "Hydrogen", false},
   			},
   		},
   	}



    http.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "application/json")


        if err := json.NewEncoder(w).Encode(questions); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    })
    fmt.Println("PORT is running on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}