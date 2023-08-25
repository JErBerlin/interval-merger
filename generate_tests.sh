#!/bin/bash

# Generate anchor points based on max_range and num_anchor_points
generate_anchor_points() {
    local max_range=$1
    local num_points=$2

    local interval_distance=$(( max_range / (num_points - 1) ))

    local anchor_points=""
    for (( i=0; i<num_points; i++ )); do
        local anchor_point=$(( interval_distance * i ))
        anchor_points+="$anchor_point "
    done

    echo "$anchor_points"
}

# Generate covering intervals based on the max_range, num_points and generated anchor points
generate_covering_intervals() {
    local max_range=$1
    local num_points=$2
    local anchor_points_output=$3

    # Convert the anchor_points_output to an array
    local anchor_points=($anchor_points_output)

    # Calculate the average distance between anchor points
    local avg_distance=$(( max_range / (num_points - 1) ))

    local intervals=""
    for anchor_point in "${anchor_points[@]}"; do
        local radius=$(( (avg_distance / 2) + (RANDOM % avg_distance) ))
        local start=$(( anchor_point - radius ))
        if (( start < 0 )); then
            start=0
        fi

        local end=$(( anchor_point + radius ))
        if (( end > max_range )); then
            end=$max_range
        fi

        intervals+="[$start,$end] "
    done

    echo "$intervals"
}

generate_test_file() {
    local range=$1
    local size=$2

    case $range in
        low) 
            max_range=100
            ;;
        mid)
            max_range=10000
            ;;
        high)
            max_range=2147483647  # Max of int32
            ;;
        *)
            echo "Invalid range!"
            exit 1
    esac

    case $size in
        xs)
            num_points=25
            ;;
        s)
            num_points=125
            ;;
        m)
            num_points=500
            ;;
        l)
            num_points=12500
            ;;
        xl)
            num_points=125000
            ;;
        *)
            echo "Invalid size!"
            exit 1
    esac

    # Set constants for anchor points and generating iterations
    local num_anchor_points=5
    local generating_iterations=5

    test_file_path="./testdata/$range.$size.txt"
    expected_file_path="./testdata/expected/$range.$size.txt"

    > $test_file_path
    > $expected_file_path

    # Generate the anchor points
    local anchors=$(generate_anchor_points $max_range $num_anchor_points)
    
    # Generate the covering intervals and write to the test file
    for (( i=0; i<$generating_iterations; i++ )); do
        local covering_intervals=$(generate_covering_intervals $max_range $num_points "$anchors")
        for interval in $covering_intervals; do
            echo "$interval" >> $test_file_path
        done
    done

    echo "[$a,$b]" >> $expected_file_path
}

# Call the function
generate_test_file $1 $2

