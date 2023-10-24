package fixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pethost/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var url = fmt.Sprintf("http://localhost:%s", config.TestPort)

type PostInput struct {
	Body     any
	URI      string
	Response any
	Token    string
}

func Post(t *testing.T, input PostInput) (statusCode int) {
	b, err := json.Marshal(input.Body)
	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, url+input.URI, bytes.NewBuffer(b))
	require.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	if input.Token != "" {
		req.Header.Set("Authorization", "Bearer "+input.Token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	require.Nil(t, err)

	statusCode = resp.StatusCode

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	if resp.Header.Get("Content-Type") != "application/json" {
		if v, ok := input.Response.(*string); ok {
			*v = string(responseBody)
		}

		return
	}

	err = json.Unmarshal(responseBody, input.Response)
	require.Nil(t, err)

	return
}

type GetInput struct {
	URI      string
	Response any
}

func Get(t *testing.T, input GetInput) (statusCode int) {
	resp, err := http.Get(url + input.URI)
	require.Nil(t, err)

	statusCode = resp.StatusCode

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	if v, ok := input.Response.(*string); ok {
		*v = string(responseBody)
		return
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return
	}

	err = json.Unmarshal(responseBody, input.Response)
	require.Nil(t, err)

	return
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

type DeleteInput struct {
	URI  string
	Body any
}

func Delete(t *testing.T, input DeleteInput) (string, int) {
	b, err := json.Marshal(input.Body)
	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodDelete, url+input.URI, bytes.NewBuffer(b))
	require.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	require.Nil(t, err)

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	return string(responseBody), resp.StatusCode
}
