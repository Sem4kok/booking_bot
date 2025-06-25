package bot

import "booking_bot/internal/models"

type Bot struct {
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

	return b
}

func GetMe() *models.User {
	u := &models.User{}

	return u
}
