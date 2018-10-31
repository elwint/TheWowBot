package main

import (
	"fmt"
	"net/http"
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
		if u.Message.Text == `/start` || strings.EqualFold(u.Message.Text, `/wow`) {
			sendWow(u.Message.Chat.ID, 0, true)
		} else if strings.EqualFold(u.Message.Text, `wow`) {
			wow <- u.Message.Chat.ID
		}
	}

	if u.InlineQuery != nil {
		q := inlineQueryResult{
			Kind:     `article`,
			ID:       `wow`,
			Title:    `wow`,
			ThumbURL: conf.InlineTumb,
		}
		q.Content.Text = conf.ASCII
		q.Content.ParseMode = `Markdown`

		call(`answerInlineQuery`, answerInlineQuery{
			ID:      u.InlineQuery.ID,
			Results: []inlineQueryResult{q},
		})
	}

	return c.NoContent(http.StatusOK)
}
