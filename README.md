# Chaotic Array Generator

## Why does this exist?

This code is basically a poor man's fuzzer for an algorithms class project. The goal of the algorithm is to partition an array of integers
into two sub-arrays that sum to the same value.

The algorithm itself is implemented in another repository. This code is designed to generate test data.

## Building

This code is implemented using Go 1.23.1. It _should_ be compatible with any Go >= 1.20, but hasn't and likely won't be tested. In order to
compile the code you can use `make`

```bash
make all
# or, with my_os equal to linux, windows, or darwin
make bin/create-array-$my_os
```

you can also compile it yourself if you're feeling feisty

```bash
go build -o bin/create-array cmd/cli.go
```

## Running

```bash
NAME:
   create-array - A Chaotic Array Generator

USAGE:
   create-array [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --trial-count value  How many rounds per exponent to create. Default is 1 (default: 1)
   --exponents value    Generate arrays of size 10^<n>. Starts with 1 up to the provided value. Default is 1 (default: 1)
   --result-file value  Path to file to save results to. Default is ./results.json (default: "results.json")
   --help, -h           show help
```

This tool is primarily designed to make a newline delimited JSON file to use as input to the real application. The results that are
generated will be placed in the `--result-file`. The `--trial-count` argument allows you to specify how many trials you'd like to run
per exponent. This is useful to be able to average your results along multiple runs for a given size. The `--exponents` flag specifies
the upper bound of your array size on a log scale.

This code will create trials for each power of 10 up to and including the `--exponents` value. For example this command

```bash
create-array --result-file output.json --trial-count 5 --exponents 3
```

Will create a `output.json` file with 15 lines. The first 5 lines will be a arrays of 10 values, then the next 5 lines will have arrays of
100 values, and finally 5 lines with arrays of 1,000 values.

## Future Plans

Maybe one day I'll fix this up. Somethings that might be fun to include are:

1. Subcommands

- It might be useful to be able to create one off arrays with particular values
