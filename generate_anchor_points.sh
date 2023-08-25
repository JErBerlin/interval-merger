#!/bin/bash

# Constants
num_anchor_points=5

generate_anchor_points() {
    local max_range=$1

    # The distance between consecutive anchor points
    local interval_distance=$((max_range / (num_anchor_points - 1)))

    local anchor_points=() # Array to store the positions of the anchor points

    # Define starting and ending points of the overarching interval
    local a=0
    local b=$max_range

    # The first anchor point is at position 'a'
    anchor_points+=($a)

    # Calculate the positions for the inner anchor points
    for (( i=1; i<$(($num_anchor_points - 1)); i++ )); do
        anchor_points+=($(($a + $interval_distance * $i)))
    done

    # The last anchor point is at position 'b'
    anchor_points+=($b)

    # Print anchor points (for debugging purpose, can be removed later)
    echo "Anchor Points: ${anchor_points[@]}"
}

# Call the function for demonstration purposes
generate_anchor_points 100  # Example for max_range = 100

