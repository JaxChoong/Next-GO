package data

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
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
	var Books []Book
	filePath := "./data/data.csv"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_,_ = reader.Read()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		Books = addToBooks(record, Books)
	}
	return Books
}

func addToBooks(record []string, Books []Book) []Book {
	var currentBook Book
	currentBook.Id = record[0]
	currentBook.Title = record[1]
	currentBook.Author = record[2]
	year, _ := strconv.Atoi(record[3])
	currentBook.Year = year
	currentBook.Description = record[4]
	currentBook.CoverImage = record[5]
	rating, _ := strconv.ParseFloat(record[6], 64)
	currentBook.Rating = rating
	pages, _ := strconv.Atoi(record[7])
	currentBook.Pages = pages
	Books = append(Books, currentBook)
	return Books
}
