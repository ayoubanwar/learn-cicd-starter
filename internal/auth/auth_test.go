package auth

import "testing"

// TestAdd tests the Add function.
func TestAdd(t *testing.T) {
    result := 2 + 3
    expected := 3

    if result != expected {
        t.Errorf("2 + 3 = %d; want %d", result, expected)
    }
}

