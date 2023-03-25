package helpers

import "testing"

func TestGenerateRandomString(t *testing.T) {
	length := 24
	s := GenerateRandomString(length)

	if len(s) != length {
		t.Errorf("expected length %v, actual %v", length, len(s))
	}
}
