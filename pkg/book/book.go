package book

import (
	"fmt"
	"gobook/pkg/author"
)

type Book struct {
	Title string
	NumberOfPages int
	Status BookStatus
	BookAuthor *author.Author
}

func New(title string, numberOfPages int, bookAuthor *author.Author) (*Book, error) {
	if err := validate(title, numberOfPages); err != nil {
		return nil, err
	}

	return &Book{
		Title: title,
		NumberOfPages: numberOfPages,
		Status: StatusAvailable,
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
