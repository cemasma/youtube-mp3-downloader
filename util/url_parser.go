package util

import (
	"net/url"
	"strings"
)

// GetVideoID parse the text and returns video id
func GetVideoID(link string, list bool) string {
	if list {
		return strings.Split(strings.Split(link, "v=")[1], "&")[0]
	}
	return strings.Split(link, "v=")[1]
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

// GetPlaylistID parses url and returns list id
func GetPlaylistID(link string) (id string) {
	parsedURL, err := url.Parse(link)
	if err == nil {
		id = parsedURL.Query().Get("list")
	}
	return
}
