package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type citiesResponse struct {
	Cities []string `json:"cities"` // Cities capitalised to export it, otherwise json encoder will ignore it.
}

func cityHandler(res http.ResponseWriter, req *http.Request) {
	cities := citiesResponse{
		Cities: []string{"Amsterdam", "Berlin", "New York", "San Francisco", "Tokyo"}}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(res).Encode(cities)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte("Hello World!"))
}

type memberResponse struct {
	Members []string `json:"members"`  //group members
}

func memberHandler(res http.ResponseWriter, req *http.Request){
	members := memberResponse{ Members: []string{"Ruby", "weifeng", "iris"}}
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(res).Encode(members)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/cities.json", cityHandler)
	http.HandleFunc("/members.json", memberHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Unable to listen on port 5000 : ", err)
	}
}
