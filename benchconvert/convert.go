package benchconvert

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type result struct {
	match       bool
	name        string
	ops         int
	nsPerOp     float64
	memoryInfo  bool
	bPerOp      float64
	allocsPerOp float64
}

// resultRegex is the regular expression used to parse the output of the go
// benchmark command. It is used to extract the name of the benchmark, the
// number of operations and the time per operation.
//
// Example:
// BenchmarkWindowIntSlice/1-1000-24          50503             24077 ns/op
// var resultRegex = regexp.MustCompile(`^Benchmark(\S*)\s+(\d+)\s+(\d+|[0-9]+.[0-9]+)\s+ns/op$`)
var resultRegex = regexp.MustCompile(`^Benchmark(\S*)\s+(\d+)\s+(\d+|[0-9]+.[0-9]+)\s+ns/op(\s+(\d+|[0-9]+.[0-9]+)\s+B/op\s+(\d+|[0-9]+.[0-9]+)\s+allocs/op)?$`)

func parseLine(line string) (res result, err error) {
	// Extract the name of the benchmark, the number of operations and the
	// time per operation.
	matches := resultRegex.FindStringSubmatch(line)
	if matches == nil {
		return result{}, nil
	}
	res.match = true

	// Extract the name of the benchmark, the number of operations and the
	// time per operation.
	res.name = matches[1]
	ops, err := strconv.Atoi(matches[2])
	if err != nil {
		return result{}, fmt.Errorf("failed to parse number of operations %s: %w", matches[2], err)
	}
	res.ops = ops
	nsPerOp, err := strconv.ParseFloat(matches[3], 64)
	if err != nil {
		return result{}, fmt.Errorf("failed to parse time per operation %s: %w", matches[3], err)
	}
	res.nsPerOp = nsPerOp

	// Extract the memory information, if available.
	if matches[4] != "" {
		res.memoryInfo = true
		bPerOp, err := strconv.ParseFloat(matches[5], 64)
		if err != nil {
			return result{}, fmt.Errorf("failed to parse bytes per operation %s: %w", matches[5], err)
		}
		res.bPerOp = bPerOp
		allocsPerOp, err := strconv.ParseFloat(matches[6], 64)
		if err != nil {
			return result{}, fmt.Errorf("failed to parse allocations per operation %s: %w", matches[6], err)
		}
		res.allocsPerOp = allocsPerOp
	}

	return res, nil
}

// ConvertToCSV converts the given stream of go benchmark output into CSV
// format. The output is written to the given writer.
func ConvertToCSV(input io.Reader, output io.Writer, quiet bool) error {
	// Read the input line by line.
	firstLine := true
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// Extract the name of the benchmark, the number of operations and the
		// time per operation.
		line := scanner.Text()
		result, err := parseLine(line)
		if err != nil {
			return fmt.Errorf("failed to parse line: %w", err)
		}
		if !result.match {
			continue
		}

		// Write the header to the output.
		if firstLine {
			header := "name,ops,ns_per_op"
			if result.memoryInfo {
				header += ",b_per_op,allocs_per_op"
			}
			if !quiet && output != os.Stdout {
				fmt.Println(header)
			}
			_, err = fmt.Fprintln(output, header)
			if err != nil {
				return fmt.Errorf("failed to write header: %w", err)
			}
			firstLine = false
		}

		// Write the benchmark result to the output.
		outputLine := fmt.Sprintf("%s,%d,%f", result.name, result.ops, result.nsPerOp)
		if result.memoryInfo {
			outputLine += fmt.Sprintf(",%f,%f", result.bPerOp, result.allocsPerOp)
		}
		if !quiet && output != os.Stdout {
			fmt.Println(outputLine)
		}
		_, err = fmt.Fprintln(output, outputLine)
		if err != nil {
			return fmt.Errorf("failed to write benchmark result: %w", err)
		}
	}

	return nil
}
