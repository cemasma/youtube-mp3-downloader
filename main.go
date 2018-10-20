package main

import (
	"flag"
	"fmt"
	"strings"
	"youtube-mp3-downloader/util"
	"youtube-mp3-downloader/youtube"
)

func main() {
	url := flag.String("url", "", "You must send it with youtube url.\nExample: youtube-mp3-downloader --url https://www.youtube.com/watch?v=t1dXdsUricU")
	playlist := flag.Bool("playlist", false, "If you want to download all of the videos from a playlist you should use this.\n"+
		"Example: youtube-mp3-downloader --url https://www.youtube.com/watch?v=80214E8FuBo&list=RDMM80214E8FuBo&start_radio=1 --playlist true")
	flag.Parse()

	if len(*url) > 0 {
		if strings.Index(*url, "list") > -1 && *playlist {
			playlistID := util.GetPlaylistID(*url)
			items := youtube.GetIDListIfExist(playlistID)
			for _, item := range items {
				util.Download("https://www.youtube.com/watch?v=" + item.Content.ID)
			}
		} else {
			videoID := util.GetVideoID(*url, false)
			if youtube.IsVideoExist(videoID) {
				util.Download(*url)
			} else {
				fmt.Println("Video doesn't exist.")
			}
		}
	}
}
