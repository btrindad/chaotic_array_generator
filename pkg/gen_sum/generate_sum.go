package gen_sum

import (
	"errors"
	"math"
	"math/rand"
)

// This function receives a target integer and returns
// a sequence of `length` integers that sum to this original target
func GenerateSum(target int, length int) ([]int, error) {
	if length == 0 {
		return nil, errors.New("Cannot generate a 0-length sequence")
	}

	seq := make([]int, 0, length)

	for i := 0; i < length-1; i++ {
		next_num := draw_number(target / length)
		target -= next_num
		seq = append(seq, next_num)
	}

	seq = append(seq, target)

	return seq, nil
}

func draw_number(target int) int {
	if target == 0 {
		return 0
	}

	value := rand.Intn(int(math.Abs(float64(target))))

	if target < 0 {
		value *= -1
	}

	return value
}
