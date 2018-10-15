package util

import (
	"testing"
)

func TestDownload(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Download Test",
			args: args{
				url: "https://www.youtube.com/watch?v=KpCnH_My-YE",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Download(tt.args.url); got != tt.want {
				t.Errorf("Download() = %v, want %v", got, tt.want)
			}
		})
	}
}
