package array

import (
	"csci_6212/project3/test_cases/pkg/gen_sum"
	"math"
	"math/rand"
	"slices"
)

func GenerateArray(target int, length int) []int {
	left_length, right_length := split_length(length)
	left_array := generate_sequence(target, left_length)
	right_array := generate_sequence(target, right_length)

	output_array := join_arrays(left_array, right_array)
	return output_array
}

// Divides a total length into two sub-lengths to generate
func split_length(length int) (int, int) {
	max_partition := int(math.Floor(float64(length) * 0.70))
	left := rand.Intn(max_partition)
	right := length - left

	return left, right
}

// A wrapper for generating a particular sub sequence
func generate_sequence(target int, length int) []int {
	seq, err := gen_sum.GenerateSum(target, length)
	if err != nil {
		// Crash if we hit an error
		panic(err)
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
