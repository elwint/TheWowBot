package main

import (
	"math/rand"
	"sync"
	"time"
)

var wow = make(chan int)
var limit = make(map[int]int)
var lock = sync.Mutex{}

func handleWow() {
	for {
		id := <-wow
		lock.Lock()
		if limit[id] < 100 {
			limit[id]++
			go sendWow(id)
		}
		lock.Unlock()
	}
}

func sendWow(id int) {
	if conf.MaxWait > 0 {
		time.Sleep(time.Duration(rand.Intn(conf.MaxWait)) * time.Second)
	}
	lock.Lock()
	limit[id]--
	lock.Unlock()

	send(id, []string{`Wow`, conf.Wow}[rand.Intn(2)])
}
