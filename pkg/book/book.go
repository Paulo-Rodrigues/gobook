package book

import "fmt"

type Book struct {
	Title string
	NumberOfPages int
	Status BookStatus
}

func New(title string, numberOfPages int) (*Book, error) {
	if err := validate(title, numberOfPages); err != nil {
		return nil, err
	}

	return &Book{
		Title: title,
		NumberOfPages: numberOfPages,
		Status: StatusAvailable,
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
