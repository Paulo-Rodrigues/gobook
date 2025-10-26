package author_test

import (
	"testing"
	"gobook/pkg/author"
)

func TestAuthorInitialization(t *testing.T) {
	a, err := author.New("John", "Doe")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if a.FirstName != "John" {
		t.Errorf("Expected first name to be 'John', got %v", a.FirstName)
	}

	if a.LastName != "Doe" {
		t.Errorf("Expected last name to be 'Doe', got %v", a.LastName)
	}
}
