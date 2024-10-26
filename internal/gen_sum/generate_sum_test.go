package gen_sum

import "testing"

func TestGenerateSum(t *testing.T) {
	tests := map[string]struct {
		target int
		length int
	}{
		"Simple Target": {
			target: 15,
			length: 4,
		},
	}

	for msg, tc := range tests {
		t.Run(msg, func(t *testing.T) {
			seq, err := GenerateSum(tc.target, tc.length)
			if err != nil {
				t.Errorf("Encountered an unexpected error generating target %d", tc.target)
			}

			if len(seq) != tc.length {
				t.Errorf("Did not generate a sequence of the correct length. Expected: %d, Got: %d", tc.length, len(seq))
			}

			seq_sum := sum_array(seq)

			if seq_sum != tc.target {
				t.Errorf("Array sum is incorrect. Expected: %d, Got: %d with array %v", tc.target, seq_sum, seq)
			}
		})
	}
}

func sum_array(seq []int) int {
	sum := 0
	for _, num := range seq {
		sum += num
	}

	return sum
}
