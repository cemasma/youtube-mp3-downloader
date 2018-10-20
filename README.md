# youtube-mp3-downloader
It provides you to download youtube videos as in mp3 format. Also you can download a list of video such as specific video.

## Installation
- <a href="https://golang.org/dl/">Install Go</a>

- <a href="https://rg3.github.io/youtube-dl/">Install youtube-dl</a> and add it to the path.<br>
- After that <a href="https://www.ffmpeg.org/">install ffmpeg</a> and add it to the path.

- And lastly run this command on your terminal.
```sh
go get -u github.com/cemasma/youtube-mp3-downloader
```

## Command-Line Usage

- If you want to download single video as in mp3 format.

```sh
youtube-mp3-downloader --url https://www.youtube.com/watch?v=t1dXdsUricU
 ```

- If you want to download whole playlist as in mp3 format.

```sh
youtube-mp3-downloader --url https://www.youtube.com/playlist?list=PLJmaGDjGiTOrgF3BFmmFX0K5eFt20dHe9 --playlist true
 ```

| Flag     | Description              |
|----------|--------------------------|
| url      | video or playlist url    |
| playlist | determine type of source |
|          |                          |