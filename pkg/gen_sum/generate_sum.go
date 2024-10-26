package gen_sum

import (
	"errors"
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
		next_num := rand.Intn(target / 2)
		target -= next_num
		seq = append(seq, next_num)
	}

	seq = append(seq, target)

	return seq, nil
}
