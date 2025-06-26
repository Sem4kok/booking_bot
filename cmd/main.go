package main

import (
	"booking_bot/internal/bapp"
	"booking_bot/internal/bot"
	"booking_bot/internal/logger"
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	sync := logger.InitLogger()
	defer sync()

	b := bot.NewBot(
		bot.WithToken("7522163268:AAFfgOfU7uiP5jVFgdweCOUKad_2ENNzk6s"),
	)

	s := server.NewServer(":8443")
	s.MustRun()

	u, err := b.GetMe()
	if err != nil {

	}
	fmt.Println(u)

	shutdown, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
	defer cancel()

	<-shutdown.Done()
	logger.Lg.Info("gracefully shutdown started")

	//s.Shutdown()
}
