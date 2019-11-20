package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"

)


func str(str string) *string {
	return &str
}

func end(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Game end - %v\n", string(body))
}

func start(w http.ResponseWriter, r *http.Request) {
	var requestData GameStartRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Game starting - %v\n", string(body))
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		log.Println(err)
		return
	}

}

func pp(val []byte) {

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, val, "", "\t")
	fmt.Println(prettyJSON.String())
}

func move(w http.ResponseWriter, r *http.Request) {
	var requestData SnakeRequest
	val, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(val, &requestData)
	responseData := MoveResponse{
		Move:  requestData.GenerateMove(),
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UP!"))
}
