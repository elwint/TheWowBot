package main

type config struct {
	Token    string `toml:"BOT_TOKEN"`
	CertFile string `toml:"CERT_FILE"`
	KeyFile  string `toml:"KEY_FILE"`
	Port     int    `toml:"PORT"`
	Route    string `toml:"ROUTE"`
	MaxWait  int    `toml:"MAX_WAIT"`
}

type update struct {
	Message struct {
		Chat struct {
			ID int `json:"id"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
}
