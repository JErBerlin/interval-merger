package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadIntervals(r io.Reader) ([]Interval, error) {
	reader := bufio.NewReader(r)

	// Read the first line to get the number of intervals
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("could not read first line: %w", err)
	}
	numIntervals, err := strconv.Atoi(string(line))
	if err != nil {
		return nil, fmt.Errorf("Error parsing the number of intervals: %w", err)
	}
	intervals := make([]Interval, 0, numIntervals) // Preallocate memory

	// Read each subsequent line as an interval
	for i := 0; i < numIntervals; i++ {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(line) == 0 {
			break
		}

		if isPrefix {
			return nil, errors.New("Encountered an unexpectedly long line for an interval.")
		}

		intervalStr := strings.Trim(strings.TrimSpace(string(line)), "[]")
		intervalPoints := strings.Split(intervalStr, ",")

		if len(intervalPoints) != 2 {
			return nil, fmt.Errorf("Invalid interval format for line: %s", string(line))
		}

		a, errA := strconv.Atoi(intervalPoints[0])
		if errA != nil {
			return nil, fmt.Errorf("Error converting string %s to integer: %w", intervalPoints[0], errA)
		}

		b, errB := strconv.Atoi(intervalPoints[1])
		if errB != nil {
			return nil, fmt.Errorf("Error converting string %s to integer: %w", intervalPoints[1], errB)
		}

		intervals = append(intervals, Interval{a: int32(a), b: int32(b)})
	}

	return intervals, nil
}

func ReadIntervalsWithValidation() ([]Interval, error) {
	// Read input intervals from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	var intervals []Interval

	// Compile the regex outside the loop
	re := regexp.MustCompile(`\[(\d+),(\d+)\](?: \[\d+,\d+\])*`)

	for scanner.Scan() {
		line := scanner.Text()

		// Validate the whole line against the regex
		if !re.MatchString(line) {
			return nil, fmt.Errorf("invalid interval format in line: %s", line)
		}

		// Extract all valid matches
		matches := re.FindAllStringSubmatch(line, -1)

		// Convert matches to Interval type
		for _, match := range matches {
			a, errA := strconv.Atoi(match[1])
			b, errB := strconv.Atoi(match[2])

			if errA != nil || errB != nil {
				return nil, fmt.Errorf("invalid numbers in interval: %s", match[0])
			}

			if a > b {
				return nil, fmt.Errorf("invalid interval: start %d is greater than end %d", a, b)
			}

			intervals = append(intervals, Interval{a: int32(a), b: int32(b)})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	return intervals, nil
}
