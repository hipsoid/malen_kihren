package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

)

var heads = []string{"beluga", "bendr", "dead", "evil", "fang", "pixel", "regular", "safe", "sand-worm", "shades", "silly", "smile", "tongue"}
var tails = []string{"block-bum", "bolt", "curled", "fat-rattle", "freckled", "hook", "pixel", "regular", "round-bum", "sharp", "skinny", "small-rattle"}


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

	responseData := GameStartResponse{
		Name:     "malen_kihren",
		HeadType: str("smile"),
		TailType: str("hook"),
	}
	b, err := json.Marshal(responseData)
	if err != nil {
		log.Println("%v", err)
		return
	}
	w.Write(b)
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
		Taunt: str(taunts[rand.Intn(len(taunts))]),
	}
	log.Printf("Move request - direction:%v - taunt: %v\n", responseData.Move, *responseData.Taunt)
	if err != nil {
		fmt.Printf("ERR: %#v\n", err)
	}
	log.Printf("%v\n", string(val))
	b, err := json.Marshal(responseData)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	w.Write(b)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UP!"))
}
