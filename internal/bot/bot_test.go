package bot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBot(t *testing.T) {
	b := NewBot(
		WithToken("123"),
		WithName("Carl"),
		WithAbout("About"),
		WithDescription("Description"),
	)

	require.Equal(t, "123", b.token, "token must be initialized")
	require.Equal(t, "Carl", b.name, "name must be initialized")
	require.Equal(t, "About", b.about, "about must be initialized")
	require.Equal(t, "Description", b.description, "description must be initialized")
	require.NotNil(t, b.app, "app must be initialized")

	b = NewBot()

	require.Equal(t, "", b.token)
	require.Equal(t, "", b.name)
	require.Equal(t, "", b.about)
	require.Equal(t, "", b.description)
	require.NotNil(t, b.app, "app must be initialized")
}
