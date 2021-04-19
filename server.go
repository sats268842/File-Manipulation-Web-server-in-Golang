package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	First     string `json:"first"`
	Last      string `json:"last"`
	Company   string `json:"company"`
	CreatedAt string `json:"created_at"`
	Country   string `json:"country"`
}

func createFile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	var customers []Customer
	err := json.NewDecoder(r.Body).Decode(&customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	JSON, _ := json.Marshal(customers)
	body, err := ioutil.ReadAll(r.Body)
	fmt.Printf(string(body))
	err = ioutil.WriteFile("customer.json", JSON, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "File added"}`))

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func readFile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	dat, err := ioutil.ReadFile("customer.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(dat)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/createFile", createFile).Methods("POST")
	router.HandleFunc("/readFile", readFile).Methods("GET")
	fmt.Printf("Starting server...............\n")
	log.Fatal(http.ListenAndServe(":8081", router))
}
