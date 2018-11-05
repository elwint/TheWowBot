package main

import (
	"math/rand"
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
			go sendWow(id)
		}
	}
}

func sendWow(id int) {
	if conf.MaxWait > 0 {
		time.Sleep(time.Duration(rand.Intn(conf.MaxWait)) * time.Second)
	}
	wow <- -1 * id

	send(id, []string{`Wow`, conf.Wow}[rand.Intn(2)])
}
