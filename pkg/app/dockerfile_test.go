package app

import "testing"

func TestImageDigest(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "valid image",
			arg:  "alpine:3.9.6",
			want: "sha256:65b3a80ebe7471beecbc090c5b2cdd0aafeaefa0715f8f12e40dc918a3a70e32",
		},
		{
			name: "invalid image",
			arg:  "alpine:invalid",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := imageDigest(tt.arg); got != tt.want {
				t.Errorf("imageDigest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizeImage(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "valid image",
			arg:  "FROM alpine:3.9.6",
			want: "FROM alpine:3.9.6@sha256:65b3a80ebe7471beecbc090c5b2cdd0aafeaefa0715f8f12e40dc918a3a70e32",
		},
		{
			name: "valid image with digest",
			arg:  "FROM alpine:3.9.6@sha256:65b3a80ebe7471beecbc090c5b2cdd0aafeaefa0715f8f12e40dc918a3a70e32",
			want: "FROM alpine:3.9.6@sha256:65b3a80ebe7471beecbc090c5b2cdd0aafeaefa0715f8f12e40dc918a3a70e32",
		},
		{
			name: "invalid image",
			arg:  "FROM alpine:invalid",
			want: "FROM alpine:invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeFrom(tt.arg); got != tt.want {
				t.Errorf("sanitizeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
