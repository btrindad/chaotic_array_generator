package array

import (
	"csci_6212/project3/test_cases/internal/util"
	"testing"
)

func TestGenerateArray(t *testing.T) {
	tests := map[string]struct {
		target int
		length int
	}{
		"Simple Array": {
			target: 30,
			length: 50,
		},
	}

	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			arr := GenerateArray(tc.target, tc.length)

			arr_sum := util.SumArray(arr)

			if arr_sum != (2 * tc.target) {
				t.Errorf("Array sum did not sum to the correct value. Expected: %d, Got: %d", (tc.target * 2), arr_sum)
			}

			if len(arr) != tc.length {
				t.Errorf("Did not generate expected number of values. Expected: %d, Got: %d", tc.length, len(arr))
			}
		})
	}
}
