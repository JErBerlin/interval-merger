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
func MergeBySort(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by interval start (a).
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].a < intervals[j].a
	})

	// Use a pointer to keep track of where we are in the merged result
	index := 0

	// Iterate through the sorted intervals and merge any overlapping ones.
	for _, interval := range intervals[1:] {
		if interval.a <= intervals[index].b {
			// Merge the intervals
			if interval.b > intervals[index].b {
				intervals[index].b = interval.b
			}
		} else {
			index++
			intervals[index] = interval
		}
	}

	return intervals[:index+1]
}
