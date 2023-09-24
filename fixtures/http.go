package fixtures

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const url = "http://localhost:3000"

type PostInput struct {
	Body map[string]string
	URI  string
}

func Post(t *testing.T, input PostInput) string {
	b, err := json.Marshal(input.Body)
	require.Nil(t, err)

	resp, err := http.Post(url+input.URI, "application/json", bytes.NewReader(b))
	require.Nil(t, err)

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	return string(responseBody)
}
