package main

import (
	"math/rand"
	"sync"
	"time"
)

var wow = make(chan int)
var limit = make(map[int]int)
var lock = sync.Mutex{}
var cancel = sync.Map{}

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
	defer func() {
		lock.Lock()
		limit[id]--
		lock.Unlock()
	}()

	var wait int
	if conf.MaxWait > 0 {
		wait = rand.Intn(conf.MaxWait)
	}

	for i := 0; i < wait; i++ {
		time.Sleep(time.Second)
		if _, ok := cancel.Load(id); ok {
			return
		}
	}
	send(id, []string{`Wow`, conf.Wow}[rand.Intn(2)])
}

func cancelWow(id int) {
	cancel.Store(id, true)
	for {
		lock.Lock()
		i := limit[id]
		lock.Unlock()
		if i == 0 {
			break
		}
	}
	cancel.Delete(id)
}
