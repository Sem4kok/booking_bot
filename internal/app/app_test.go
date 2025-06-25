package app

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewApp(t *testing.T) {
	a := NewApp("123")

	require.NotNil(t, a.c)
	require.Equal(t, "https://api.telegram.org/bot123/", a.plate)
}

func TestWithPlate(t *testing.T) {
	a := NewApp("123")

	require.Equal(t, "https://api.telegram.org/bot123/test", a.withPlate("test"))
}
