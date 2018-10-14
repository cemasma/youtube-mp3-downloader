package main

import (
	"flag"
)

func main() {
	url := flag.String("url", "", "You must send it with youtube url.\nExample: ymp3d --url https://www.youtube.com/watch?v=t1dXdsUricU")
	flag.Parse()

	if len(*url) > 0 {

	}
}
