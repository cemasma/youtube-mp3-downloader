package util

import "testing"

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
