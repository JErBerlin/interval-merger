/*
   Interval Merger

   Description:
   This program computes the set union of overlapping integer intervals provided as string input.

   Assumptions and technical decisions:
   - Intervals are provided in the form [a,b], where a and b are integers, representable as int32.
   - The set of intervals to merge are provided from stdin, one interval for each line>
   -- the first line of the input will indicate the maximum lines with intervals to be read
   -- the input can be ended with an EOF (for redirected input files)
   -- the indicated lines will be read or until an EOF is found, without causing an execution error
   - The start (a) of an interval is always less than or equal to its end (b).
   - The number of intervals in the input are representable by the positiv part of an int32

	Stragety:
   - The program uses a sorting and merging algorithm to compute the union of intervals.

   Example Input: [25,30] [2,19] [14,23] [4,8]
   Example Output: [2,23] [25,30]
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// Read input intervals from stdin.

	// Obs: we assume input correctly formatted for speed
	// for validation, use ReadIntervalsWithValidation
	intervals, err := ReadIntervals(os.Stdin)
	if err != nil {
		log.Fatal("Error reading intervals: ", err)
	}

	// Merge overlapping intervals using sort and merge.
	merged := mergeBySort(intervals)

	// Output to stdout using Stringer of Intervals
	fmt.Println(Intervals(merged))
}
