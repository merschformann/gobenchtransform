# gobenchtransform

Transforms `go test` benchmark output to CSV.

## Usage

```bash
go test -bench . | gobenchtransform > bench.csv
```

Output:

```csv
name,ops,ns_per_op
Slice-24,20372328,55.510000
Map-24,86512916,11.610000
```

Please refer to the help manual (`gobenchtransform -h`) for further information. The example above is located in the `examples/` directory.

## Installation

```bash
go install github.com/merschformann/gobenchtransform@latest
```

## Why?

This is a small tool I wrote for converting `go test` benchmark output into a format (CSV) that I could easily use for plotting and other downstream purposes. I shared it just in case anyone else finds it useful.
