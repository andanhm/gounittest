package tbt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response tells gounittest
type Response struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  bool   `json:"status"`
}

// Curl request given url using http.Get in go (golang)
func Curl(url string) (*Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%s", response.Status)
	}
	r := new(Response)
	err = json.NewDecoder(response.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
