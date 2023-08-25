package main

import (
	"fmt"
	"os"
	"testing"
)

const (
	TestFilesPath     = "./testdata/"
	ExpectedFilesPath = "./testdata/expected/"
)

func TestEndToEnd(t *testing.T) {
	// Categories of range and size for test files
	rangePatterns := []string{"low", "mid", "high"}
	sizePatterns := []string{"xs", "s", "m", "l", "xl"}

	// Iterate over range and size patterns
	for _, rangePattern := range rangePatterns {
		for _, sizePattern := range sizePatterns {
			testFileName := fmt.Sprintf("%s.%s.txt", rangePattern, sizePattern)
			expectedFileName := testFileName // since they share the same name

			testFilePath := fmt.Sprintf("%s%s", TestFilesPath, testFileName)
			expectedFilePath := fmt.Sprintf("%s%s", ExpectedFilesPath, expectedFileName)

			// Check if both files exist
			if _, err := os.Stat(testFilePath); os.IsNotExist(err) {
				// Test file doesn't exist, skip to next combination
				continue
			}
			if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
				t.Fatalf("Expected file %s doesn't exist for test file %s", expectedFilePath, testFilePath)
			}

			// Read test file and compute merged intervals
			testFile, _ := os.Open(testFilePath)
			intervals, err := ReadIntervals(testFile)
			if err != nil {
				t.Fatalf("Error reading intervals from %s: %s", testFilePath, err)
			}
			merged := MergeBySort(intervals)

			// Read expected result
			expectedFile, _ := os.Open(expectedFilePath)
			expectedIntervals, err := ReadIntervals(expectedFile)
			if err != nil {
				t.Fatalf("Error reading intervals from %s: %s", expectedFilePath, err)
			}

			// Compare results
			if !equal(merged, expectedIntervals) {
				t.Fatalf("For file %s, expected %v but got %v", testFileName, expectedIntervals, merged)
			}
		}
	}
}

// Helper function to compare two slices of intervals
func equal(a, b []Interval) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
