package services

import "testing"

func TestMD5HashService_HashString(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{
		{"12345", "827ccb0eea8a706c4c34a16891f84e7b"},
		{"Admin12345", "e66055e8e308770492a44bf16e875127"},
	}

	svc := NewMD5HashService()

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			actual := svc.HashString(test.input)

			if actual != test.output {
				t.Errorf("expected: %v, actual: %v", test.output, actual)
			}
		})
	}
}
