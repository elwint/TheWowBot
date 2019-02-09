package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/elwint/router"
)

var conf = config{}

func main() {
	if _, err := toml.DecodeFile(`config.toml`, &conf); err != nil {
		panic(err)
	}
	go handleWow()

	r := router.New()
	r.POST(conf.Route, postUpdate)
	panic(r.StartTLS(fmt.Sprintf(`:%d`, conf.Port), conf.CertFile, conf.KeyFile))
}

func postUpdate(c *router.Context, u update) error {
	if u.Message != nil {
		message(u.Message.Chat.ID, u.Message.Text)
	}
	if u.InlineQuery != nil {
		inlineQuery(u.InlineQuery.ID)
	}

	return c.NoContent(http.StatusOK)
}

func message(id int, text string) {
	text = strings.TrimSuffix(text, `@`+conf.Username)

	if text == `/start` || strings.EqualFold(text, `/wow`) {
		send(id, conf.Wow)
	} else if strings.EqualFold(text, `/rip`) {
		send(id, conf.RIP)
	} else if strings.EqualFold(text, `/cancel`) {
		cancelWow(id)
	} else if matched, _ := regexp.MatchString(`(?i)illuminati|triangle|driehoek`, text); matched {
		send(id, []string{conf.Illuminati, conf.StickerIlluminati}[rand.Intn(2)])
	} else if matched, _ := regexp.MatchString(`(?i)wo+?w`, text); matched {
		wow <- id
	}
}

func inlineQuery(id string) {
	call(`answerInlineQuery`, answerInlineQuery{
		ID: id,
		Results: []inlineQueryResult{
			result(`Illuminati Sticker`, conf.StickerIlluminati),
			result(`Wow`, conf.Wow),
			result(`RIP`, conf.RIP),
			result(`Illuminati`, conf.Illuminati),
		},
	})
}
