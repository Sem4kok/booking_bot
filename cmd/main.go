package main

import (
	"booking_bot/internal/bot"
	"fmt"
)

func main() {
	b := bot.NewBot(
		bot.WithToken("7522163268:AAGPflKoOywa-pNSS0jZGKCEUlmsTyUzrPE"),
	)

	u, err := b.GetMe()
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
}
