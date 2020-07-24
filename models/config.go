package models

type Config struct {
	BotToken string `json:"bot_token"`
	WebHook  string `json:"web_hook"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
