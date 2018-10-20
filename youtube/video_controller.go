package youtube

import (
	"encoding/json"
	"log"
	"net/http"
	"youtube-mp3-downloader/util"
)

// ContentDetail contains video id for playlist videos
type ContentDetail struct {
	ID string `json:"videoId"`
}

// Item is data structure for store video details
type Item struct {
	ID      string        `json:"id"`
	Kind    string        `json:"kind"`
	Content ContentDetail `json:"contentDetails"`
}

type videos struct {
	Items         []Item `json:"items"`
	NextPageToken string `json:"nextPageToken"`
}

// apiKey is Youtube api v3 public key
const apiKey string = "AIzaSyAMP1Mnjpo0Vl8l_HhOw5nKP8RzQ67EhKE"

// IsVideoExist checks video existence
func IsVideoExist(id string) bool {
	resp, err := http.Get("https://www.googleapis.com/youtube/v3/videos?part=id&id=" + id + "&key=" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var video videos
	json.NewDecoder(resp.Body).Decode(&video)
	if len(video.Items) > 0 && video.Items[0].ID == id {
		return true
	}

	return false
}

// GetIDListIfExist is checks the playlist and returns video ids
func GetIDListIfExist(id string) []Item {
	items := []Item{}
	isThereNextPageToken := true

	queryURL := "https://www.googleapis.com/youtube/v3/playlistItems?part=contentDetails&playlistId=" + id + "&key=" + apiKey
	for isThereNextPageToken == true {
		resp, err := http.Get(queryURL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var video videos
		json.NewDecoder(resp.Body).Decode(&video)

		if len(video.Items) == 0 {
			return items
		} else if video.NextPageToken != "" {
			queryURL = util.AddNextPageToken(queryURL, video.NextPageToken)
		} else {
			isThereNextPageToken = false
		}

		items = append(items, video.Items...)

	}

	return items
}
