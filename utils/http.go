package utils

import (
	"io"
	"net/http"
)

func GetFileContentFromUrl(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
