package run_trials

import (
	"csci_6212/project3/test_cases/pkg/array"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func check_args(trials int, exponent int) error {
	if trials < 1 {
		msg := fmt.Sprintf("Invalid number of trials: %d", trials)
		return errors.New(msg)
	}

	if exponent < 1 {
		msg := fmt.Sprintf("Invalid exponent: %d", exponent)
		return errors.New(msg)
	}

	return nil
}

func GenerateData(output_file string, trials int, exponent int) error {
	err := check_args(trials, exponent)
	if err != nil {
		return err
	}

	file, err := os.Create(output_file)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	for pow := 1; pow <= exponent; pow++ {
		for i := 0; i < trials; i++ {
			round, err := get_trial(int(math.Pow10(pow)))
			if err != nil {
				return err
			}

			encoder.Encode(round)
		}
	}

	return nil
}

type Trial struct {
	Sum   int
	Array []int
}

// A helper function to make an array of a certain length
func get_trial(length int) (Trial, error) {
	sum := rand_sum(length)
	min_value := rand_min()
	arr, err := array.GenerateArray(sum, length, min_value)
	if err != nil {
		return Trial{}, err
	}

	round := Trial{
		Sum:   sum,
		Array: arr,
	}

	return round, nil
}

// A tiny helper function to allow me to tune the frequency of
// arrays with negative numbers
func rand_min() int {
	roll := rand.Intn(10)
	if roll <= 2 {
		// Only return a negative number ~30% of the time
		return (rand.Intn(25) * -1)
	} else {
		return 0
	}
}

// A helper function to find an appropriate sum based on the
// length of the array
func rand_sum(length int) int {
	upper := (length * 10) - 1
	lower := (length / 2) + length
	return rand.Intn(upper-lower) + lower
}
