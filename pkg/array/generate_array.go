package array

import (
	"csci_6212/project3/test_cases/pkg/gen_sum"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"slices"
)

func GenerateArray(target int, length int, min_value int) ([]int, error) {
	err := check_args(target, length, min_value)
	if err != nil {
		return nil, err
	}

	left_length, right_length := split_length(length)
	left_array := generate_sequence(target, left_length, 0)
	right_array := generate_sequence(target, right_length, min_value)

	output_array := join_arrays(left_array, right_array)
	return output_array, nil
}

// Divides a total length into two sub-lengths to generate
func split_length(length int) (int, int) {
	max_partition := int(math.Floor(float64(length) * 0.70))
	left := rand.Intn(max_partition) + 1
	right := length - left

	return left, right
}

// A wrapper for generating a particular sub sequence
func generate_sequence(target int, length int, min_value int) []int {
	shifted_target := (min_value * -1 * length) + target

	seq, err := gen_sum.GenerateSum(shifted_target, length)
	if err != nil {
		// Crash if we hit an error
		panic(err)
	}

	// This check isn't necessary, but potentially saves time
	if min_value != 0 {
		for i := 0; i < length; i++ {
			seq[i] = seq[i] + min_value
		}
	}

	return seq
}

// Join two given arrays and shuffle them for entropy
func join_arrays(left []int, right []int) []int {
	output_array := slices.Concat(left, right)

	// Use a built-in to implement Fisher-Yates Shuffle
	rand.Shuffle(len(output_array), func(i, j int) {
		output_array[i], output_array[j] = output_array[j], output_array[i]
	})

	return output_array
}

func check_args(target int, length int, min_value int) error {
	if !(length > 2) {
		msg := fmt.Sprintf("Invalid array length: %d", length)
		return errors.New(msg)
	}

	return nil
}
