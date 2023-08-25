package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadIntervals(r io.Reader) ([]Interval, error) {
	reader := bufio.NewReader(r)
	var intervals []Interval

	// Read the first line to get the number of intervals
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("could not read first line: %w", err)
	}
	numIntervals, err := strconv.Atoi(string(line))
	if err != nil {
		return nil, fmt.Errorf("Error parsing the number of intervals: %w", err)
	}

	// Read each subsequent line as an interval
	for i := 0; i < numIntervals; i++ {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Println("found EOF in read loop:", err)
				break // end of the file reached
			}
			log.Println("found not EOF error in read loop:", err)
			return nil, err
		}

		if len(line) == 0 {
			break // no content line, end of manual input
		}

		if isPrefix {
			// The line is too long, and we are only getting a fragment.
			// This shouldn't happen as intervals should fit on one line.
			return nil, errors.New("Encountered an unexpectedly long line for an interval.")
		}

		part := strings.TrimSpace(string(line))
		intervalStr := strings.Trim(part, "[]")
		intervalPoints := strings.Split(intervalStr, ",")

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

	if len(intervals) == 0 {
		return nil, errors.New("read intervals: could not read any interval")
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

	return intervals, scanner.Err()
}
