package utils

import (
	"io"
	"net/http"
)

func ApiCall(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseBytes, _ := io.ReadAll(response.Body)
	return responseBytes, nil
}
