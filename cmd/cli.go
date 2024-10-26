package main

import (
	run_trials "csci_6212/project3/test_cases/pkg"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "create-array",
		Usage: "A Chaotic Array Generator",
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

			return run_trials.GenerateData(file, trials, exp)
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Unexpected Error Ocurred", "msg", err.Error())
		os.Exit(1)
	}
}
