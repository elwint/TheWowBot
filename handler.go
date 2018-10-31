package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var wow = make(chan int)

func handleWow() {
	chat := make(map[int]int)
	for {
		id := <-wow
		if id < 0 {
			chat[-1*id]--
			continue
		}

		if chat[id] < 1000 {
			chat[id]++
			go sendWow(id, conf.MaxWait)
		}
	}
}

func sendWow(id int, maxWait int) {
	if maxWait > 0 {
		time.Sleep(time.Duration(rand.Intn(maxWait)) * time.Second)
	}
	wow <- -1 * id

	call(`sendMessage`, sendMessage{
		ID:   id,
		Text: `wow`,
	})
}

func call(method string, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
		return
	}

	url := fmt.Sprintf(`https://api.telegram.org/bot%s/%s`, conf.Token, method)
	resp, err := http.Post(url, `application/json`, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	r, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("%d: %s\n", resp.StatusCode, r)
	}
}
