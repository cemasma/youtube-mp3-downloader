package youtube

import (
	"reflect"
	"testing"
)

func TestIsVideoExist(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsVideoExist 1",
			args: args{
				url: "b_bJQgZdjzo",
			},
			want: true,
		},
		{
			name: "TestIsVideoExist 2",
			args: args{
				url: "wqww",
			},
			want: false,
		},
		{
			name: "TestIsVideoExist with ID in Playlist",
			args: args{
				url: "qXN15uh4DLU",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := IsVideoExist(tt.args.url); got != tt.want {
			t.Errorf("%q. IsVideoExist() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetIDListIfExist(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want []Item
	}{
		{
			name: "Playlist Test",
			args: args{
				url: "PLJmaGDjGiTOrgF3BFmmFX0K5eFt20dHe9",
			},
			want: []Item{
				{"UExKbWFHRGpHaVRPcmdGM0JGbW1GWDBLNWVGdDIwZEhlOS41NkI0NEY2RDEwNTU3Q0M2", "youtube#playlistItem"},
				{"UExKbWFHRGpHaVRPcmdGM0JGbW1GWDBLNWVGdDIwZEhlOS4yODlGNEE0NkRGMEEzMEQy", "youtube#playlistItem"},
			},
		},
		{
			name: "Wrong Url Test",
			args: args{
				url: "ww",
			},
			want: []Item{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIDListIfExist(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIDListIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
