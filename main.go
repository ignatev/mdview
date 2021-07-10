package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type inputFile struct {
	filepath string
}

func main() {
	fileData, err := getFileData()
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(fileData)

}

func getFileData() (inputFile, error) {
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("A filepath argument is required")
	}

	fileLocation := flag.Arg(0)

	return inputFile{fileLocation}, nil
}
