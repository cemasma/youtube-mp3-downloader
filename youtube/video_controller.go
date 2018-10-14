package youtube

import (
	"encoding/json"
	"log"
	"net/http"
)

type item struct {
	ID   string `json:"id"`
	Etag string `json:"etag"`
	Kind string `json:"kind"`
}

type videos struct {
	Items []item `json:"items"`
}

// apiKey is Youtube api v3 public key
const apiKey string = "AIzaSyAMP1Mnjpo0Vl8l_HhOw5nKP8RzQ67EhKE"

// IsVideoExist checks video existence
func IsVideoExist(url string) bool {
	resp, err := http.Get("https://www.googleapis.com/youtube/v3/videos?part=id&id=" + url + "&key=" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var video videos
	json.NewDecoder(resp.Body).Decode(&video)
	if len(video.Items) > 0 && video.Items[0].ID == url {
		return true
	}

	return false
}
