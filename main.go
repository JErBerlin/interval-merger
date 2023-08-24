/*
   Interval Merger

   Description:
   This program computes the set union of overlapping integer intervals provided as string input.

   Assumptions:
   - Intervals are provided in the form [a,b], where a and b are integers, representable as int32.
   - The set of intervals to merge are provided from stdin in string form, as a line of space-separated intervals.
   - The start (a) of an interval is always less than or equal to its end (b).
   - The number of intervals in the input are representable by the positiv part of an int32.

	Stragety:
   - The program uses a sorting and merging algorithm to compute the union of intervals.

   Example Input: [25,30] [2,19] [14,23] [4,8]
   Example Output: [2,23] [25,30]
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Interval represents a numeric range with a start (a) and end (b).
type Interval struct {
	a int32
	b int32
}

func main() {
	// Read input intervals from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	var intervals []Interval
type Intervals []Interval

	for scanner.Scan() {
		line := scanner.Text()
func (i Intervals) String() string {

		// TODO: validate input format
		parts := strings.Split(line, " ")
	var s strings.Builder

		// Convert each interval in string format to Interval type.
		for _, part := range parts {
			// discard empty parts (extra spaces or newline)
			if part == "" {
				continue
			}
	// concatenate strings efficiently with strings builder
	for _, interval := range i {
		s.WriteString(fmt.Sprintf("[%d,%d] ", interval.a, interval.b))
	}
	s.WriteString("\n")

			intervalContent := strings.Trim(part, "[]")
			intervalPoints := strings.Split(intervalContent, ",")
	return s.String()
}

			// TODO: validate integer numbers / handle conversion error
			a, _ := strconv.Atoi(intervalPoints[0])
			b, _ := strconv.Atoi(intervalPoints[1])
			intervals = append(intervals, Interval{a: int32(a), b: int32(b)})
		}
	}

	// Merge overlapping intervals.
	merged := mergeBySort(intervals)

	// Print the merged intervals.
	for _, interval := range merged {
		fmt.Printf("[%d,%d] ", interval.a, interval.b)
	}
	fmt.Println()
}

// mergeBySort takes a list of intervals, sorts them by start (a), and merges any overlapping intervals.
func mergeBySort(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by start (a).
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].a < intervals[j].a
	})

	// Initialize the merged intervals list with the first interval.
	merged := []Interval{intervals[0]}

	// Iterate through the sorted intervals and merge any overlapping ones.
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]

		// If the current interval overlaps with the last merged interval, merge them.
		if interval.a <= last.b {
			if interval.b > last.b {
				last.b = interval.b
				// TODO: check next line, necessary?
				merged[len(merged)-1] = last
			}
		} else {
			merged = append(merged, interval)
		}
	}

	return merged
}
