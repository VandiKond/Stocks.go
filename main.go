package main

import (
	"fmt"
	"time"

	"github.com/VandiKond/Stocks.go/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/robfig/cron/v3"
)

func job() {
	fmt.Println("This job runs every minute.", time.Now())
}

func main() {
	c := cron.New()
	_, err := c.AddFunc("* * * * *", job)
	if err != nil {
		panic(err)
	}
	c.Start()

	bot, err := tgbotapi.NewBotAPI(Tocken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			messages.Reply(update, update.Message.Text, nil, bot)
		}
	}

	select {}
}
