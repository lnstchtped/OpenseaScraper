package internal

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func DownloadImage(url string) ([]byte, string, error) {
	if !strings.Contains(url, "png") {
		return nil, "", nil
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header = http.Header{
		"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0"},
		"accept":     {"image/png,*/*"},
		"referer":    {"https://opensea.io/"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	if strings.HasPrefix(string(body), "<") {
		return nil, "", fmt.Errorf("invalid response")
	}

	filepath := strings.Split(strings.Split(url, "files/")[1], "?")[0]

	return body, filepath, nil
}
