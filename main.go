package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ReadinessRequest struct {
	endpoint string
}

type ReadinessResponse struct {
	endpoint string
	code     int
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bastion Host")
}
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	var command ReadinessRequest
	err := json.NewDecoder(r.Body).Decode(&command)

	//body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, command)
	/*
		res, err := http.Get(command.endpoint)
		if err != nil {
			log.Fatalln(command.endpoint)
			log.Fatalln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Fprint(w, body)

		}
	*/
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/readiness", ReadinessHandler)
	log.Fatal(http.ListenAndServe(":46802", r))
}
