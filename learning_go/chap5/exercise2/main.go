package main

import (
	"fmt"
	"os"
)

func fileLen(filename string) (int, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	fstat, err := fd.Stat()
	if err != nil {
		return 0, err
	}

	return int(fstat.Size()), nil
}

func main() {
	l, err := fileLen("/home/sha2ks/Downloads/Createaser-168248630116211-15_54_39.mp4")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(l)
}
