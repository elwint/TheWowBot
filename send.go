package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func send(id int, text string) {
	if text == conf.StickerIlluminati {
		call(`sendSticker`, sendSticker{
			ID:      id,
			Sticker: text,
		})
		return
	}

	call(`sendMessage`, sendMessage{
		ID:        id,
		Text:      text,
		ParseMode: `Markdown`,
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
	defer resp.Body.Close() // nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		r, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%d: %s\n", resp.StatusCode, r)
	}
}
