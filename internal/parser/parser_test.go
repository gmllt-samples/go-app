package parser

import (
	"testing"
)

func TestParseStatus(t *testing.T) {
	tests := []string{"200", "200-204", "200,201,404"}
	for _, input := range tests {
		_, err := ParseStatus(input)
		if err != nil {
			t.Errorf("ParseStatus failed for input %q: %v", input, err)
		}
	}
}

func TestParseSize(t *testing.T) {
	tests := []string{"100", "1K", "5K-10K", "1K,2K,3K"}
	for _, input := range tests {
		size, err := ParseSize(input)
		if err != nil {
			t.Errorf("ParseSize failed for input %q: %v", input, err)
		}
		if size <= 0 {
			t.Errorf("ParseSize returned non-positive value for %q", input)
		}
	}
}

func TestParseDuration(t *testing.T) {
	tests := []string{"500ms", "1s-2s", "100ms,200ms,300ms"}
	for _, input := range tests {
		d, err := ParseDuration(input)
		if err != nil {
			t.Errorf("ParseDuration failed for input %q: %v", input, err)
		}
		if d <= 0 {
			t.Errorf("ParseDuration returned non-positive duration for %q", input)
		}
	}
}
