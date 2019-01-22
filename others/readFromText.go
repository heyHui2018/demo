package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"io"
)

func main() {
	file, err := os.Open("fake.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		data, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		info := string(data)
		list := strings.Split(info, ",")
		if len(list) == 3 {
			fmt.Println(list)
		}
	}
}
