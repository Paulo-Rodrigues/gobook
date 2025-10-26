package book_test

import (
	"testing"
	"gobook/pkg/book"
)

func TestStatuses(t *testing.T) {
	if book.StatusAvailable.String() != "available" {
		t.Errorf("Expected StatusAvailable to be 'available', got '%s'", book.StatusAvailable.String())
	}

	if book.StatusUnavailable.String() != "unavailable" {
		t.Errorf("Expected StatusUnavailable to be 'unavailable', got '%s'", book.StatusUnavailable.String())
	}

	if book.StatusRented.String() != "rented" {
		t.Errorf("Expected StatusRented to be 'rented', got '%s'", book.StatusRented.String())
	}
}
