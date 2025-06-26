package main

import (
	"booking_bot/internal/bot"
	"booking_bot/internal/logger"
	"fmt"
)

func main() {
	sync := logger.InitLogger()
	defer sync()

	b := bot.NewBot(
		bot.WithToken("7522163268:AAGPflKoOywa-pNSS0jZGKCEUlmsTyUzrPE"),
	)

	u, err := b.GetMe()
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
}
