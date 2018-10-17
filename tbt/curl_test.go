package tbt

import (
	"reflect"
	"testing"
)

func TestCurl(t *testing.T) {

	tests := []struct {
		name    string
		url     string
		want    *Response
		wantErr bool
	}{
		{
			"url exist",
			"https://www.andanhm.me/gounittest.json",
			&Response{
				Name:    "gounittest",
				Version: "v1.0.0",
				Status:  true,
			},
			false,
		},
		{
			"url not exist",
			"https://www.andanhm.me/not_exist.json",
			nil,
			true,
		},
		{
			"url provided invalid",
			"andanhm.me/not_exist.json",
			nil,
			true,
		},
		{
			"expected json parser error",
			"https://www.andanhm.me/invalid.json",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := Curl(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Curl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(response, tt.want) {
				t.Errorf("Curl() error = %v, want %v", err, tt.want)
				return
			}
		})
	}
}
