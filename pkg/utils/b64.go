package utils

import (
	"encoding/base64"
)

func EncodeB64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeB64(data string) (string, error) {
	s, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
