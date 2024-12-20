package main

import (
	"Teleram_GO/errors"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

type Config struct {
	Token string `json:"token"`
}

func loadJson(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(errors.ErrorOpenFile)
	}
	defer file.Close()
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("%s:\n", errors.ErrorDecoderFile, err)
	}
	return &config, nil
}

func main() {
	config, err := loadJson("config/config.json")
	if err != nil {
		log.Fatalf("%s: %v", errors.ErrorLoadConfiguration, err)
	}

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(errors.ErrrorRunBot, err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}
}
