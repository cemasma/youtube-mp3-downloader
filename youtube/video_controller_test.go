package youtube

import "testing"

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
	}
	for _, tt := range tests {
		if got := IsVideoExist(tt.args.url); got != tt.want {
			t.Errorf("%q. IsVideoExist() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
