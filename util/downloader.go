package util

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Download function is download videos via youtube-dl
func Download(url string) bool {
	cmd := exec.Command("youtube-dl", "--extract-audio", "--audio-format", "mp3", url)
	log.Printf("Downloading...")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("An error occured: %v", err)
	}
	if strings.Index(string(out), "[download] 100%") > -1 {
		fmt.Printf("\nDownload is done.")
		return true
	}
	return false
}
