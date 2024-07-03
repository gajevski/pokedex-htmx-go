package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func ExtractID(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", err
	}
	path := u.Path
	parts := strings.Split(path, "/")
	if len(parts) > 1 {
		number := parts[len(parts)-2]
		return number, nil
	} else {
		fmt.Println("URL path format is incorrect")
	}
	return "", nil
}
