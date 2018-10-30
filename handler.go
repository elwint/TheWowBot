package main

import (
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

	resp, err := http.Get(fmt.Sprintf(`https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=wow`, conf.Token, id))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("%d: %s\n", resp.StatusCode, b)
	}
}
