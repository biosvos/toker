package toker_test

import (
	_ "embed"
	"testing"
	"time"

	"github.com/biosvos/toker"
	"github.com/stretchr/testify/require"
)

//go:embed test/data/test.dev.crt
var devCRT []byte

//go:embed test/data/test.dev.key
var devKey []byte

func TestNewPrivateToker(t *testing.T) {
	t.Parallel()
	_, err := toker.NewPrivateToker(devKey)

	require.NoError(t, err)
}

func TestPrivateToker_Generate(t *testing.T) {
	t.Parallel()
	tok, _ := toker.NewPrivateToker(devKey)

	_, err := tok.Generate(time.Now().Add(15*time.Minute), nil)

	require.NoError(t, err)
}

type Payload struct {
	Name string
}

func generateToken(t *testing.T, expired time.Time, name string) string {
	t.Helper()
	private, err := toker.NewPrivateToker(devKey)
	require.NoError(t, err)

	token, err := private.Generate(expired, &Payload{
		Name: name,
	})
	require.NoError(t, err)

	return token
}

func TestNewPublicToker(t *testing.T) {
	t.Parallel()
	_, err := toker.NewPublicToker(devCRT)

	require.NoError(t, err)
}

func TestPublicToker_ParseToken(t *testing.T) {
	t.Parallel()
	tok, _ := toker.NewPublicToker(devCRT)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		token := generateToken(t, time.Now().Add(15*time.Minute), "ASDF")
		var payload Payload

		err := tok.Parse(token, &payload)

		require.NoError(t, err)
		require.Equal(t, "ASDF", payload.Name)
	})
	t.Run("expired token", func(t *testing.T) {
		t.Parallel()
		token := generateToken(t, time.Now().Add(-15*time.Minute), "ASDF")
		var payload Payload

		err := tok.Parse(token, &payload)

		require.ErrorIs(t, err, toker.ErrExpiredToken)
	})
}
