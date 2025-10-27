package book_test

import (
	"testing"
	"gobook/pkg/book"
)

func TestBookFormats(t *testing.T) {
	if book.FormatHardcover.String() != "hardcover" {
		t.Errorf("Expected FormatHardcover to be 'hardcover', got '%s'", book.FormatHardcover.String())
	}

	if book.FormatPaperback.String() != "paperback" {
		t.Errorf("Expected FormatPaperback to be 'paperback', got '%s'", book.FormatPaperback.String())
	}

	if book.FormatEbook.String() != "ebook" {
		t.Errorf("Expected FormatEbook to be 'ebook', got '%s'", book.FormatEbook.String())
	}
}

