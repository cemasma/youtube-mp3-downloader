package util

import "strings"

// GetVideoID parse the text and returns video id
func GetVideoID(url string, list bool) string {
	if list {
		return strings.Split(strings.Split(url, "v=")[1], "&")[0]
	}
	return strings.Split(url, "v=")[1]
}

// AddNextPageToken is adds or changes the pageToken
func AddNextPageToken(queryURL, pageToken string) (newURL string) {
	if strings.Index(queryURL, "&pageToken") > -1 {
		newURL = strings.Replace(queryURL, strings.Split(queryURL, "&pageToken=")[1], pageToken, 1)
	} else {
		newURL = queryURL + "&pageToken=" + pageToken
	}
	return
}
