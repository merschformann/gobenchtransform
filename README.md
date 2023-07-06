# gobenchtransform

Transforms go test bench output to CSV.

## Usage

```bash
go test -bench . -benchmem | gobenchtransform > bench.csv
```

## Installation

```bash
go install github.com/merschformann/gobenchtransform@latest
```

## Help

```bash
$ gobenchtransform --help
Transform Go benchmark results into a format that can be used by other tools.

Usage:
  gobenchtransform [flags]

Flags:
  -h, --help            help for gobenchtransform
  -i, --input string    input file (default is stdin)
  -o, --output string   output file (default is stdout)
  -q, --quiet           suppress repeating output to stdout
```
