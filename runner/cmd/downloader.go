package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var (
	ERR_BAD_STATUS = fmt.Errorf("bad status code")
	ERR_NOT_AUTHED = fmt.Errorf("not authenticated")
)

func DownloadInput(year int, day int, token string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: token,
	})

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%w: %d", ERR_BAD_STATUS, resp.StatusCode)
	}

	buf := bytes.Buffer{}
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

type Sample struct {
	Input  string `json:"input"`
	Output int    `json:"output"`
}

func ExtractSamples(day int, part int) ([]Sample, error) {
	return nil, nil
}
