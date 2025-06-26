package bapp

import (
	"booking_bot/internal/bot"
	"booking_bot/internal/logger"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"net/http"
)

type Bapp struct {
	s    *fiber.App
	port string

	b *bot.Bot

	lg *zap.Logger
}

func NewBapp(port string, bot *bot.Bot) *Bapp {
	return &Bapp{
		s: fiber.New(),
		lg: logger.Lg.With(
			zap.String("service", "server"),
		),
		b:    bot,
		port: port,
	}
}

func (b *Bapp) MustRun() {
	go func() {
		b.lg.Info(
			"server starting",
			zap.String("port", b.port),
		)

		b.s.Post("/update", b.Update)

		b.lg.Fatal(
			"server errors:",
			zap.String("error:", b.s.Listen(b.port).Error()),
		)
	}()
}

func (b *Bapp) Update(c fiber.Ctx) error {
	b.lg.Info("update has been handled")

	c.JSON(http.StatusOK)
	return nil
}

func (b *Bapp) Shutdown() {
	b.s.Shutdown()
}
