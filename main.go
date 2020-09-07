package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ReadinessRequest struct {
	Endpoint string `json:"endpoint"`
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
	//s := string(body)
	//json.Unmarshal([]byte(s), &command)
	//fmt.Println(s)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := http.Get(command.Endpoint)

	if err != nil {
		//log.Fatalln(command.Endpoint)
		//log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]byte("500 - Something bad happened!"))
	} else {
		fmt.Println(res)
		w.WriteHeader(http.StatusOK)
	}

	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Fprint(w, body)
	//fmt.Println(body)

}
func main() {
	fmt.Println("Go server on 46802 port")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/readiness", ReadinessHandler)
	log.Fatal(http.ListenAndServe(":46802", r))
}
