package main

import (
	"fmt"
	"sort"
	"strings"
)

// Interval represents a numeric range with a start (a) and end (b).
type Interval struct {
	a int32
	b int32
}

type Intervals []Interval

func (i Intervals) String() string {

	var s strings.Builder

	// concatenate strings efficiently with strings builder
	for _, interval := range i {
		s.WriteString(fmt.Sprintf("[%d,%d] ", interval.a, interval.b))
	}
	s.WriteString("\n")

	return s.String()
}

// mergeBySort takes a list of intervals, sorts them by start (a), and merges any overlapping intervals.
// TODO: use type Intervals as argument, or define function as its method
// TODO: make function exported if needed for tests or new package
func mergeBySort(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by interval start (a).
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
				// TODO: check if next line can be optimised
				merged[len(merged)-1] = last
			}
		} else {
			merged = append(merged, interval)
		}
	}

	return merged
}
