package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerJsonError(t *testing.T) {
	msg := "Hello json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello json"}`), result)
}
