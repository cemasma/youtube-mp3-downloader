package util

import "strings"

// GetVideoID parse the text and returns video id
func GetVideoID(url string, list bool) string {
	if list {
		return strings.Split(strings.Split(url, "v=")[1], "&")[0]
	}
	return strings.Split(url, "v=")[1]
}
