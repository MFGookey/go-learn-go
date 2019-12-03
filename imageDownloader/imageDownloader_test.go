package main

import (
	"testing"
)

type fakeImage struct {
	Filename string
}

func (f fakeImage) Download() string {
	return f.Filename
}

func TestDoAThing(t *testing.T) {
	type args struct {
		i Downloader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Has Filename", args: args{i: fakeImage{Filename: "asdf"}}, want: "asdf", wantErr: false},
		{name: "Has Filename", args: args{i: fakeImage{Filename: ""}}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DoAThing(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoAThing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoAThing() = %v, want %v", got, tt.want)
			}
		})
	}
}
