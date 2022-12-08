package helper

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type MsgType struct {
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
}

func GetMessage() string {

	var data string
	fd := OpenFile()
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("error reading file %s", err)
			break
		}
		data = fmt.Sprintf("%s \n %s", data, line)
	}
	return data
}

func (msg MsgType) saveMessage() {

	fd, err := os.OpenFile("./log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	if _, err := fd.WriteString(fmt.Sprintf("%s : %s\n", msg.Timestamp, msg.Data)); err != nil {
		log.Fatal(err)
	}
}
