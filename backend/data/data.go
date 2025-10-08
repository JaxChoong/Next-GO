package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Book struct {
	Id          string
	Title       string
	Author      string
	Year        int
	Description string
	CoverImage  string
	Rating      float64
	Pages       int
}

func Data() []Book {
	// var Books []Book
	filePath := "./data/data.csv"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		addToBooks(record)
	}
	return nil
}

func addToBooks(record []string) {
	for _, word := range record {
		print(word)
	}
	fmt.Println()
}
