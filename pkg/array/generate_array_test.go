package array

import (
	"csci_6212/project3/test_cases/internal/util"
	"testing"
)

func TestGenerateArray(t *testing.T) {
	tests := map[string]struct {
		target    int
		length    int
		min_value int
		wantError bool
	}{
		"Simple Array": {
			target:    30,
			length:    50,
			min_value: 0,
			wantError: false,
		},

		"Array with negative value": {
			target:    30,
			length:    60,
			min_value: -5,
			wantError: false,
		},

		"Larger Array": {
			target:    843,
			length:    130,
			min_value: -84,
			wantError: false,
		},

		"Array with min value larger than target": {
			target:    30,
			length:    75,
			min_value: -60,
			wantError: false,
		},

		"Array with a positive min value": {
			target:    30,
			length:    75,
			min_value: 10,
			wantError: false,
		},

		"Array of 0": {
			target:    30,
			length:    0,
			min_value: 0,
			wantError: true,
		},

		"Negative Length": {
			target:    10,
			length:    -1,
			min_value: 0,
			wantError: true,
		},
	}

	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			arr, err := GenerateArray(tc.target, tc.length, tc.min_value)

			if tc.wantError && err == nil {
				t.Fatalf("Did not receive the error we expected")
			}

			if !tc.wantError {
				if err != nil {
					t.Fatalf("Received unexpected error %s", err.Error())
				}

				arr_sum := util.SumArray(arr)

				t.Logf("Debug Test <%s>: Sum => %d, Array => %v", msg, arr_sum, arr)

				if arr_sum != (2 * tc.target) {
					t.Errorf("Array sum did not sum to the correct value. Expected: %d, Got: %d", (tc.target * 2), arr_sum)
				}

				if len(arr) != tc.length {
					t.Errorf("Did not generate expected number of values. Expected: %d, Got: %d", tc.length, len(arr))
				}
			}
		})
	}
}
