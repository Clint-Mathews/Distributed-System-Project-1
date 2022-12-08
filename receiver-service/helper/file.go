package helper

import (
	"log"
	"os"
)

func CreateFile() {
	_, err := os.Create("./log.log")
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteFile() {
	err := os.Remove("./log.log")
	if err != nil {
		log.Fatal(err)
	}
}

func OpenFile() *os.File {
	fd, err := os.Open("./log.log")
	if err != nil {
		log.Fatal(err)
	}
	return fd
}
