package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}
