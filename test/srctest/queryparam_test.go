package srctest

import (
	"testing"

	"github.com/kenf1/delegator/src/auth"
)

func TestSanitizeQueryParam(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		//valid
		{"abc123", "abc123", false},
		{"ABC-123-xyz", "ABC-123-xyz", false},
		{"123-456", "123-456", false},

		//invalid
		{"abc_123", "", true},
		{"123@abc", "", true},
		{"test!", "", true},
		{"hello world", "", true},
		{"<script>", "", true},
	}

	for _, tt := range tests {
		result, err := auth.SanitizeQueryParam(tt.input)
		if tt.expectError && err == nil {
			t.Errorf("Expected error for input '%s', but got none", tt.input)
		}
		if !tt.expectError && err != nil {
			t.Errorf("Did not expect error for input '%s', but got: %v", tt.input, err)
		}
		if result != tt.expected {
			t.Errorf("Input: '%s' | Expected: '%s', Got: '%s'", tt.input, tt.expected, result)
		}
	}
}
