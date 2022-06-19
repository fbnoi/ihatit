package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"

	"fbnoi.com/ihatit/internal"
)

func main() {
	var mapper internal.Mapper
	file, err := os.Open("../xml/mapper.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(data, &mapper)
	if err != nil {
		panic(err)
	}

	log.Print(mapper)
}
