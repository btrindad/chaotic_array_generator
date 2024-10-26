package main

import (
	"csci_6212/project3/test_cases/pkg/array"
	"encoding/json"
	"log/slog"
	"math/rand"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "trial-count",
				Value: 1,
				Usage: "How many rounds per exponent to create. Default is 1",
			},

			&cli.IntFlag{
				Name:  "exponents",
				Value: 1,
				Usage: "Generate arrays of size 10^<n>. Starts with 1 up to the provided value. Default is 1",
			},

			&cli.PathFlag{
				Name:  "result-file",
				Value: "results.json",
				Usage: "Path to file to save results to. Default is ./results.json",
			},
		},
		Action: func(ctx *cli.Context) error {
			file := ctx.Path("result-file")
			trials := ctx.Int("trial-count")
			exp := ctx.Int("exponents")

			return generate_data(file, trials, exp)
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Unexpected Error Ocurred", "msg", err.Error())
		os.Exit(1)
	}
}

type Trial struct {
	Sum   int
	Array []int
}

func generate_data(output_file string, trials int, exponent int) error {
	file, err := os.OpenFile(output_file, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	for pow := 0; pow < exponent; pow++ {
		for i := 0; i < trials; i++ {
			// round, err := get_trial(int(math.Pow10(pow)))
			round, err := get_trial(30)
			if err != nil {
				return err
			}

			encoder.Encode(round)
		}
	}

	return nil
}

// A helper function to make an array of a certain length
func get_trial(length int) (Trial, error) {
	sum := 42 // rand.Int()
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
	if roll <= 3 {
		// Only return a negative number ~30% of the time
		return (rand.Intn(25) * -1)
	} else {
		return 0
	}
}