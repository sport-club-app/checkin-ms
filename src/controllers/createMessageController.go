package controllers

import (
	servicesKafka "checkin/src/services/kafka"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func CreateMessageController(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body) // response body is []byte
	//fmt.Println(string(body))
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	servicesKafka.Producer(body)
	json.NewEncoder(w).Encode(result)
}
