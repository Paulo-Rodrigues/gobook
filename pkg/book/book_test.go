package book_test

import (
	"testing"
	"gobook/pkg/book"
)

func TestInitializeBook(t *testing.T) {
	b, err := book.New("The Go Programming Language", 256)

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
	_, err := book.New("", -10)

	if err == nil {
		t.Fatalf("Expected error for invalid book parameters, got nil")
	}

	if err.Error() != "Title cannot be empty" && err.Error() != "Number of pages cannot be negative" {
		t.Errorf("Unexpected error message: %v", err)
	}
}
