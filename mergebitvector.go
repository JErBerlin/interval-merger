package main

// Using declarations from mergesort.go

// BitVector to represent our bit-level intervals.
type BitVector []uint64

// Set the bits for a given interval.
func (bv *BitVector) SetInterval(start, end int32) {
	for i := start; i <= end; i++ {
		(*bv)[i/64] |= 1 << (i % 64)
	}
}

// Extract merged intervals from the bit vector.
func (bv *BitVector) ExtractIntervals() []Interval {
	var intervals []Interval
	var currentInterval Interval
	inInterval := false

	for i := 0; i < len(*bv)*64; i++ {
		if (*bv)[i/64]&(1<<(i%64)) != 0 {
			if !inInterval {
				currentInterval.a = int32(i)
				inInterval = true
			}
		} else if inInterval {
			currentInterval.b = int32(i - 1)
			intervals = append(intervals, currentInterval)
			inInterval = false
		}
	}

	// If we ended in the middle of an interval.
	if inInterval {
		currentInterval.b = int32(len(*bv)*64 - 1)
		intervals = append(intervals, currentInterval)
	}

	return intervals
}

func MergeByBitVector(intervals []Interval) []Interval {
	// Determine the range.
	var maxVal int32 = 0
	for _, interval := range intervals {
		if interval.b > maxVal {
			maxVal = interval.b
		}
	}

	// Initialize the bit vector.
	bv := BitVector(make([]uint64, maxVal/64+1))

	// Set bits for each interval.
	for _, interval := range intervals {
		bv.SetInterval(interval.a, interval.b)
	}

	// Extract merged intervals.
	return bv.ExtractIntervals()
}
