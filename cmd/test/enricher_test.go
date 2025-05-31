package main_test

import (
	"net/http"
	"strings"
	"testing"
)

func Test_Enrich(t *testing.T) {

	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Create 1",
			input: `{"author": "Author1", "quote": "This is a quote 1"}`,
		},
		{
			name:  "Create 2",
			input: `{"author": "Author 2", "quote": "This is a quote 2"}`,
		},
		{
			name:  "Create 3",
			input: `{"author": "Author 3", "quote": "This is a quote 3"}`,
		},
		{
			name:  "men 1",
			input: `{"author": "men", "quote": "This is a my quote 1"}`,
		},
		{
			name:  "men  2",
			input: `{"author": "men", "quote": "This is a my quote 2"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Post("http://localhost:8080/quotes", "application/json", strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("Failed to create quote: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Expected status OK, got %s", resp.Status)
			}
		})
	}
}
