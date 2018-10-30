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
	if strings.EqualFold(u.Message.Text, `/start`) {
		sendWow(u.Message.Chat.ID, 0)
	} else if strings.EqualFold(u.Message.Text, `wow`) {
		wow <- u.Message.Chat.ID
	}

	return c.NoContent(http.StatusOK)
}