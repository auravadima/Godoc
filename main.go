package main

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
    "os"
	"encoding/json"
	"fmt"
)

type Config struct {
    TelegramBotToken string
}

func GetToken() string{
	file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    configuration := Config{}
    err := decoder.Decode(&configuration)
    if err != nil {
       log.Panic(err)
    }
    return configuration.TelegramBotToken
}

func main() {
	bot, err := tgbotapi.NewBotAPI(GetToken())
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bot.GetMe())
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates { 
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		fmt.Println(update.Message.Text)
		text := getDoc(update.Message.Text)
		msg.Text = text
		bot.Send(msg)
	}
}