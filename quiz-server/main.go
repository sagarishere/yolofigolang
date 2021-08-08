package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type Response struct {
	Answers []string `json:"Answers"`
}

var Responses []Response

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/add", newAnswer)
	myRouter.HandleFunc("/all", showAnswers)

	log.Print("Development server started on localhost:10000")
	log.Print("It's also available in local network at ", GetOutboundIP(), ":10000")
	err := http.ListenAndServe(":10000", myRouter)
	if err != nil {
		log.Fatal(err)
	}
}

func showAnswers(w http.ResponseWriter, r *http.Request) {
	for _, res := range Responses {
		json.NewEncoder(w).Encode(res)
	}
}

func newAnswer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var response Response
	json.Unmarshal(reqBody, &response)
	if response.Answers != nil {
		Responses = append(Responses, response)
		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Fprintf(w, "You tried to POST empty or wrong request. Check it's format.\n\nCorrect format: {\"Answers\":[\"Name\",\"1\",\"2\",\"3\",\"4\",\"5\"]}")
	}
}

func main() {
	Responses = []Response{
		{Answers: []string{"Name", "1", "2", "3", "4", "5", "6", "7"}},
		{Answers: []string{"John", "Yes", "No", "What", "Nope", "Simon"}},
		{Answers: []string{"Joana", "No", "Yes", "Not yet", "Yes!", "Sebastian", "-1", "Done!"}},
	}
	handleRequests()
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
