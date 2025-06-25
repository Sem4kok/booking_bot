package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"booking_bot/internal/models"
)

type App struct {
	plate string
	c     *http.Client
}

func NewApp(token string) *App {
	return &App{
		plate: fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		c:     &http.Client{},
	}
}

func (a *App) withPlate(method string) string {
	return a.plate + method
}

func (a *App) GetMe() (*models.User, error) {
	const op = "App.GetMe"

	resp, err := a.c.Get(
		a.withPlate("getMe"),
	)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}
	defer resp.Body.Close()

	var u *models.User
	if err := json.NewDecoder(resp.Body).Decode(u); err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	return u, nil
}
