package handlers

import "testing"

func TestIsSuccessStatus(t *testing.T) {
	for _, test := range []struct {
		statusCode int
		expected   bool
	}{
		{statusCode: 199, expected: false},
		{statusCode: 200, expected: true},
		{statusCode: 204, expected: true},
		{statusCode: 300, expected: false},
	} {
		if actual := isSuccessStatus(test.statusCode); actual != test.expected {
			t.Errorf("isSuccessStatus(%d) = %v, want %v", test.statusCode, actual, test.expected)
		}
	}
}