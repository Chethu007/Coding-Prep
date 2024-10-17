package main

import (
	"encoding/base64"
	"strings"
)

func FileIdEncode(id string) string {
	b := base64.StdEncoding.EncodeToString([]byte(id))
	b = strings.ReplaceAll(b, "+", ".")
	b = strings.ReplaceAll(b, "/", "_")
	b = strings.ReplaceAll(b, "=", "-")
	return b
}

func FileIdDecode(encoded string) (string, error) {
	b := encoded
	b = strings.ReplaceAll(b, ".", "+")
	b = strings.ReplaceAll(b, "_", "/")
	b = strings.ReplaceAll(b, "-", "=")
	r, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func main() {
	println("ID:", FileIdEncode("43135427")) //NDMxMzU0Mjc-
	id, _ := FileIdDecode("NDMxMzU0Mjc-")
	println("ID:", id)
}
