package book_test

import (
	"testing"
	"gobook/pkg/author"
	"gobook/pkg/book"
)

func TestInitializeBook(t *testing.T) {
	b, err := book.New("The Go Programming Language", 256, nil, book.FormatPaperback)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if b.Title != "The Go Programming Language" {
		t.Errorf("Expected title to be 'The Go Programming Language', got %v", b.Title)
	}

	if b.NumberOfPages != 256 {
		t.Errorf("Expected number of pages to be 256, got %v", b.NumberOfPages)
	}

	if b.Status != book.StatusAvailable {
		t.Errorf("Expected status to be 'Available', got %v", b.Status)
	}
}

func TestValidations(t *testing.T) {
	_, err := book.New("", -10, nil, book.FormatPaperback)

	if err == nil {
		t.Fatalf("Expected error for invalid book parameters, got nil")
	}

	if err.Error() != "Title cannot be empty" && err.Error() != "Number of pages cannot be negative" {
		t.Errorf("Unexpected error message: %v", err)
	}
}

func TestBookWithAuthor(t *testing.T) {
	a, _ := author.New("Jhon", "Doe")
	b, _ := book.New("Learning Go", 300, a, book.FormatEbook)

	if b.BookAuthor == nil {
		t.Errorf("Expected book to have an author, got nil")
	}

	expected := "Jhon Doe"

	if b.BookAuthor.FullName() != expected {
		t.Errorf("Expected author name to be %v, got %v", expected, b.BookAuthor.FullName())
	}
}

func TestBookWithFormat(t *testing.T) {
	a, _ := author.New("Robert", "Martin")
	b, _ := book.New("Clean Code", 464, a, book.FormatHardcover)

	if b.Format != book.FormatHardcover {
		t.Errorf("Expected format to be %v, got %v", book.FormatHardcover, b.Format)
	}

	if b.Format.String() != "hardcover" {
		t.Errorf("Expected format string to be 'hardcover', got %v", b.Format.String())
	}
}

func TestBookDefaultFormat(t *testing.T) {
	b, _ := book.New("Default Format Book", 200, nil, book.FormatPaperback)

	if b.Format != book.FormatPaperback {
		t.Errorf("Expected format to be %v, got %v", book.FormatPaperback, b.Format)
	}
}

func TestBookFormatValidation(t *testing.T) {
	a, _ := author.New("Martin", "Fowler")
	
	formats := []book.BookFormat{
		book.FormatHardcover,
		book.FormatPaperback,
		book.FormatEbook,
	}

	for _, format := range formats {
		b, err := book.New("Refactoring", 448, a, format)
		
		if err != nil {
			t.Errorf("Expected no error for valid format %v, got %v", format, err)
		}

		if b.Format != format {
			t.Errorf("Expected format to be %v, got %v", format, b.Format)
		}
	}
}
