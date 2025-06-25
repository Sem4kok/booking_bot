package bot

import (
	"booking_bot/internal/app"
	"booking_bot/internal/models"
)

type Bot struct {
	app         *app.App
	token       string
	name        string
	about       string
	description string
}

type Option func(*Bot)

func WithToken(token string) Option {
	return func(b *Bot) {
		b.token = token
	}
}

func WithName(name string) Option {
	return func(b *Bot) {
		b.name = name
	}
}

func WithAbout(about string) Option {
	return func(b *Bot) {
		b.about = about
	}
}

func WithDescription(description string) Option {
	return func(b *Bot) {
		b.description = description
	}
}

func NewBot(opts ...Option) *Bot {
	b := &Bot{}

	for _, opt := range opts {
		opt(b)
	}

	b.app = app.NewApp(b.token)

	return b
}

func (b *Bot) GetMe() (*models.User, error) {
	u, err := b.app.GetMe()
	if err != nil {
		return nil, err
	}

	return u, nil
}
