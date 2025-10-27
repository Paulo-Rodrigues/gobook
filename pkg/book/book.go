package book

import (
	"fmt"
	"gobook/pkg/author"
)

type Book struct {
	Title string
	NumberOfPages int
	Status BookStatus
	Format BookFormat
	BookAuthor *author.Author
}

func New(title string, numberOfPages int, bookAuthor *author.Author, format BookFormat) (*Book, error) {
	if err := validate(title, numberOfPages); err != nil {
		return nil, err
	}

	return &Book{
		Title: title,
		NumberOfPages: numberOfPages,
		Status: StatusAvailable,
		Format: format,
		BookAuthor: bookAuthor,
	}, nil
}

func validate(title string, numberOfPages int) error {
	if title == "" {
		return fmt.Errorf("Title cannot be empty")
	}

	if numberOfPages < 0 {
		return fmt.Errorf("Number of pages cannot be negative")
	}

	return nil
}
