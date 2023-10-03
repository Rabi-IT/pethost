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

type RawPostInput struct {
	Body any
	URI  string
}

func RawPost(t *testing.T, input RawPostInput) string {
	b, err := json.Marshal(input.Body)
	require.Nil(t, err)

	resp, err := http.Post(url+input.URI, "application/json", bytes.NewReader(b))
	require.Nil(t, err)

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	return string(responseBody)
}

type GetInput struct {
	URI      string
	Response any
}

func Get(t *testing.T, input GetInput) {
	resp, err := http.Get(url + input.URI)
	require.Nil(t, err)

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	err = json.Unmarshal(responseBody, input.Response)
	require.Nil(t, err)
}

type PatchInput struct {
	URI  string
	Body any
}

func Patch(t *testing.T, input PatchInput) string {
	b, err := json.Marshal(input.Body)
	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodPatch, url+input.URI, bytes.NewBuffer(b))
	require.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	return string(responseBody)
}
