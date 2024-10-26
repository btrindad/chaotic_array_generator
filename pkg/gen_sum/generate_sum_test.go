package gen_sum

import (
	"csci_6212/project3/test_cases/internal/util"
	"testing"
)

func TestGenerateSum(t *testing.T) {
	tests := map[string]struct {
		target int
		length int
	}{
		"Simple Target": {
			target: 15,
			length: 4,
		},
		"Longer Target": {
			target: 15,
			length: 20,
		},
		"Large Target": {
			target: 400,
			length: 50,
		},
	}

	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			seq, err := GenerateSum(tc.target, tc.length)

			t.Logf("TestCase %s: Sequence %v", msg, seq)

			if err != nil {
				t.Errorf("Encountered an unexpected error generating target %d", tc.target)
			}

			if len(seq) != tc.length {
				t.Errorf("Did not generate a sequence of the correct length. Expected: %d, Got: %d", tc.length, len(seq))
			}

			seq_sum := util.SumArray(seq)

			if seq_sum != tc.target {
				t.Errorf("Array sum is incorrect. Expected: %d, Got: %d with array %v", tc.target, seq_sum, seq)
			}
		})
	}
}
