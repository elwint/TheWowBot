package main

type config struct {
	Token      string `toml:"BOT_TOKEN"`
	CertFile   string `toml:"CERT_FILE"`
	KeyFile    string `toml:"KEY_FILE"`
	Port       int    `toml:"PORT"`
	Route      string `toml:"ROUTE"`
	MaxWait    int    `toml:"MAX_WAIT"`
	ASCII      string `toml:"ASCII_WOW"`
	InlineTumb string `toml:"INLINE_TUMB"`
}

type update struct {
	Message *struct {
		Chat struct {
			ID int `json:"id"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
	InlineQuery *struct {
		ID string `json:"id"`
	} `json:"inline_query"`
}

type sendMessage struct {
	ID        int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type answerInlineQuery struct {
	ID        string              `json:"inline_query_id"`
	Results   []inlineQueryResult `json:"results"`
	CacheTime int                 `json:"cache_time,omitempty"`
}

type inlineQueryResult struct {
	Kind    string `json:"type"`
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content struct {
		Text      string `json:"message_text"`
		ParseMode string `json:"parse_mode,omitempty"`
	} `json:"input_message_content"`
	Description string `json:"description,omitempty"`
	ThumbURL    string `json:"thumb_url,omitempty"`
	ThumbWidth  int    `json:"thumb_width,omitempty"`
	ThumbHeight int    `json:"thumb_height,omitempty"`
}
