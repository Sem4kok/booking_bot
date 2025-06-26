package app

import (
	"booking_bot/internal/logger"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"

	"booking_bot/internal/models"
)

type App struct {
	plate string
	lg    *zap.Logger
	c     *http.Client
}

func NewApp(token string) *App {
	return &App{
		plate: fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		lg: logger.Lg.With(
			zap.String("service", "app"),
		),
		c: &http.Client{},
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

	var r models.APIResponse[models.User]
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	if !r.OK {
		return nil, fmt.Errorf("%s : API returns not ok", op)
	}

	return &r.Result, nil
}
