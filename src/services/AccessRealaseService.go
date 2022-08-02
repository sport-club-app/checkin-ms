package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ResponsePartners struct {
	Start      string `json:"start"`
	DueDate    string `json:"dueDate"`
	PartnerId  int    `json:"partnerId"`
	ModalityId int    `json:"modalityId"`
	IsActive   bool   `json:"isActive"`
}

func AccessRealaseService(id int, token []string) ResponsePartners {

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	requestURL := fmt.Sprintf("http://10.0.0.172:3001/v1/contracts/%d", id)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("token", token[1])

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}
	var result ResponsePartners
	json.Unmarshal([]byte(body), &result)

	return result
}