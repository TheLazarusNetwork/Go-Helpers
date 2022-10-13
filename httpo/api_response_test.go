package httpo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestPayload struct {
	Name string
}

func Test_NewSuccessResponse(t *testing.T) {
	payload := TestPayload{
		Name: "Om",
	}
	p := NewSuccessResponse(200, "user fetch success", payload)
	require.Equal(t, payload.Name, p.Payload.Name)
}
