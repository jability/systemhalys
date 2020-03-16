package systemhalys

import (
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	c := Config{
		data: map[string]string{
			"FULL":     "ok!",
			"SPLIT_UP": "ok!",
		},
	}

	tests := []struct {
		input   string
		output  string
		success bool
	}{
		{"FULL", "ok!", true},
		{"SPLIT_UP", "ok!", true},
		{"NOT_FOUND", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := c.Get(tt.input)
			if tt.success == (err != nil) {
				t.Errorf("Expected '%v', got error '%v'", tt.output, err)
			}
			if got != tt.output {
				t.Errorf("Expected '%v' for key '%v', got '%v'", tt.output, tt.input, got)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	content := `INDEX1 OK
INDEX_2 OK
   INDEX_3 OK
#INDEX_4 NOK
 INDEX-5 OK`

	c := Load(strings.NewReader(content))

	tests := []struct {
		input   string
		output  string
		success bool
	}{
		{"INDEX1", "OK", true},
		{"INDEX_2", "OK", true},
		{"INDEX_3", "OK", true},
		{"INDEX_4", "", false},
		{"INDEX-5", "OK", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := c.Get(tt.input)
			if tt.success == (err != nil) {
				t.Errorf("Expected '%v', got error '%v'", tt.output, err)
			}
			if got != tt.output {
				t.Errorf("Expected '%v' for key '%v', got '%v'", tt.output, tt.input, got)
			}
		})
	}
}
