package main

import (
	"errors"
	"fmt"

	slow "github.com/jpreese/slowimage"
)

func main() {
	var i slow.Image
	name, err := DoAThing(i)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(name)
	}
}

// Downloader an interface for things that download
type Downloader interface {
	Download() string
}

// DoAThing checks an image's filename and if it is blank returns an error.  Otherwise it returns the filename.
func DoAThing(i Downloader) (string, error) {
	fileName := i.Download()

	if fileName == "" {
		return "", errors.New("hey you gave me a blank filename here")
	}

	return fileName, nil
}
