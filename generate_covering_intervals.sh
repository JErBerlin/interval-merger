#!/usr/bin/env bash
# File: generate_covering_intervals.sh

# Constants
num_anchor_points=5
generating_iterations=5
max_range=100
num_points=25

# Assume this is the output of generate_anchor_points or is a known set
generate_anchor_points_output="0 25 50 75 100"

# Function to generate covering intervals
generate_covering_intervals() {
    local max_range=$1
    local num_points=$2

    # Convert the generate_anchor_points_output to an array
    anchor_points=($generate_anchor_points_output)

    # Calculate the average distance between anchor points
    local avg_distance=$(( max_range / (num_anchor_points - 1) ))

    # Outer loop for generating_iterations
    for (( iteration=0; iteration<$generating_iterations; iteration++ )); do
        # Inner loop for each anchor point
        for (( i=0; i<$num_anchor_points; i++ )); do
            # Extract the current anchor point from the array
            local anchor_point=${anchor_points[$i]}

            # Calculate a random radius around the anchor point, considering the average distance
            local radius=$(( (avg_distance / 2) + (RANDOM % avg_distance) ))

            # Calculate the start and end of the interval
            local start=$(( anchor_point - radius ))
            if (( start < 0 )); then
                start=0
            fi

            local end=$(( anchor_point + radius ))
            if (( end > max_range )); then
                end=$max_range
            fi

            echo "[$start,$end]"
        done
    done
}

# Call the function with provided constants
generate_covering_intervals $max_range $num_points

