package util

import (
	"testing"
)

func TestGetVideoID(t *testing.T) {
	type args struct {
		url  string
		list bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test With List",
			args: args{
				url:  "https://www.youtube.com/watch?v=aJOTlE1K90k&list=PLx0sYbCqOb8TBPRdmBHs5Iftvv9TPboYG",
				list: true,
			},
			want: "aJOTlE1K90k",
		},
		{
			name: "Test Without list",
			args: args{
				url:  "https://www.youtube.com/watch?v=aJOTlE1K90k",
				list: false,
			},
			want: "aJOTlE1K90k",
		},
	}
	for _, tt := range tests {
		if got := GetVideoID(tt.args.url, tt.args.list); got != tt.want {
			t.Errorf("%q. GetVideoID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAddNextPageToken(t *testing.T) {
	type args struct {
		queryURL  string
		pageToken string
	}
	tests := []struct {
		name       string
		args       args
		wantNewURL string
	}{
		{
			name: "Test Without Page Token",
			args: args{
				queryURL:  "https://www.googleapis.com/youtube/v3/playlistItems?part=contentDetails&playlistId=asd&key=qqq",
				pageToken: "test",
			},
			wantNewURL: "https://www.googleapis.com/youtube/v3/playlistItems?part=contentDetails&playlistId=asd&key=qqq&pageToken=test",
		},
		{
			name: "Test With Token",
			args: args{
				queryURL:  "https://www.googleapis.com/youtube/v3/playlistItems?part=contentDetails&playlistId=asd&key=qqq&pageToken=test",
				pageToken: "testwithtoken",
			},
			wantNewURL: "https://www.googleapis.com/youtube/v3/playlistItems?part=contentDetails&playlistId=asd&key=qqq&pageToken=testwithtoken",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewURL := AddNextPageToken(tt.args.queryURL, tt.args.pageToken); gotNewURL != tt.wantNewURL {
				t.Errorf("AddNextPageToken() = %v, want %v", gotNewURL, tt.wantNewURL)
			}
		})
	}
}
