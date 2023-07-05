package benchconvert

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

// resultRegex is the regular expression used to parse the output of the go
// benchmark command. It is used to extract the name of the benchmark, the
// number of operations and the time per operation.
//
// Example:
// BenchmarkWindowIntSlice/1-1000-24          50503             24077 ns/op
var resultRegex = regexp.MustCompile(`^Benchmark(\S*)\s+(\d+)\s+(\d+)\s+ns/op$`)

func parseLine(line string) (name string, ops int, nsPerOp int, match bool, err error) {
	// Extract the name of the benchmark, the number of operations and the
	// time per operation.
	matches := resultRegex.FindStringSubmatch(line)
	if matches == nil {
		return "", 0, 0, false, nil
	}

	// Extract the name of the benchmark, the number of operations and the
	// time per operation.
	name = matches[1]
	ops, err = strconv.Atoi(matches[2])
	if err != nil {
		return "", 0, 0, false, fmt.Errorf("failed to parse number of operations - %s: %w", matches[2], err)
	}
	nsPerOp, err = strconv.Atoi(matches[3])
	if err != nil {
		return "", 0, 0, false, fmt.Errorf("failed to parse time per operation - %s: %w", matches[3], err)
	}

	return name, ops, nsPerOp, true, nil
}

// ConvertToCSV converts the given stream of go benchmark output into CSV
// format. The output is written to the given writer.
func ConvertToCSV(input io.Reader, output io.Writer) error {
	// Read the input line by line.
	firstLine := true
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// Extract the name of the benchmark, the number of operations and the
		// time per operation.
		line := scanner.Text()
		name, ops, nsPerOp, match, err := parseLine(line)
		if err != nil {
			return fmt.Errorf("failed to parse line: %w", err)
		}
		if !match {
			continue
		}

		// Write the header to the output.
		if firstLine {
			_, err = fmt.Fprintln(output, "name,ops,ns_per_op")
			if err != nil {
				return fmt.Errorf("failed to write header: %w", err)
			}
			firstLine = false
		}

		// Write the benchmark result to the output.
		_, err = fmt.Fprintf(output, "%s,%d,%d\n", name, ops, nsPerOp)
		if err != nil {
			return fmt.Errorf("failed to write benchmark result: %w", err)
		}
	}

	return nil
}
