#!/bin/bash

# TODO: not critical issues: 
#       - check use or not of local variables across the script       
#       - for high numbers and small sizes make incompatible parameters
#       - for high numbers and high sizes number of intervals created less than expected. It produces EOF in logs.

# Generate anchor points based on max_range and num_anchor_points
generate_anchor_points() {
    local max_range=$1
    local num_anchor_points=$2

    local distance=$(( max_range / (num_anchor_points - 1) ))

    local anchor_points=""
    for (( i=0; i<num_anchor_points; i++ )); do
        local anchor_point=$(( distance * i ))
        anchor_points+="$anchor_point "
    done

    echo "$anchor_points"
}

# Generate covering intervals based on the max_range, num_intervals and generated anchor points
generate_covering_intervals() {
    local max_range=$1
    local num_anchor_points=$2
    local anchor_points_str=$3

    # Convert the anchor_points_str to an array
    local anchor_points=($anchor_points_str)

    # Calculate the average distance between anchor points
    local avg_distance=$(( max_range / (num_anchor_points - 1) ))

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

ensure_directories_exist() {
    # Check if "./testdata/" exists
    if [ ! -d "./testdata/" ]; then
        mkdir -p "./testdata/"
    fi

    # Check if "./testdata/expected/" exists
    if [ ! -d "./testdata/expected/" ]; then
        mkdir -p "./testdata/expected/"
    fi
}

generate_test_file() {
    local range=$1
    local size=$2

    case $range in
        low) 
            max_range=100
            num_anchor_points=$(( $max_range / 20 ))
            ;;
        mid)
            max_range=10000
            num_anchor_points=$(( $max_range / 100 ))
            ;;
        high)
            max_range=2147483647  # Max of int32
            num_anchor_points=$(( $max_range / 1000000 ))
            ;;
        *)
            echo "Invalid range!"
            exit 1
    esac

    case $size in
        xs)
            num_intervals=25
            ;;
        s)
            num_intervals=125
            ;;
        m)
            num_intervals=500
            ;;
        l)
            num_intervals=12500
            ;;
        xl)
            num_intervals=125000
            ;;
        # Warning: xxl can take few minutes to generate
        xxl)
            num_intervals=1000000
            ;;
        *)
            echo "Invalid size!"
            exit 1
    esac

    # Number intervals effectively written should be approximately the required num_intervals 
    generating_iterations=$(( $num_intervals / $num_anchor_points ))

    # Ensure the directories exist, creating them if necessary
    ensure_directories_exist
    
    test_file_path="./testdata/$range.$size.txt"
    expected_file_path="./testdata/expected/$range.$size.txt"

    # Write the total number of intervals at the beginning of test and expected result files
    # For the expected result file it `is always 1, by construction of the intervals
    echo $num_intervals > $test_file_path
    echo 1 > $expected_file_path
    
    # Generate the anchor points
    local anchors=$(generate_anchor_points $max_range $num_anchor_points)
    
    # Generate the covering intervals and write to the test file
    for (( i=0; i<$generating_iterations; i++ )); do
        local covering_intervals=$(generate_covering_intervals $max_range $num_anchor_points "$anchors")
        for interval in $covering_intervals; do
            echo "$interval" >> $test_file_path
        done
    done

    echo "[0,$max_range]" >> $expected_file_path
}

# Check arguments are present when calling generate_tests
if [ -z "$1" ] || [ -z "$2" ]; then
    echo "Usage: $0 <range> <size>, where range in [low, mid, high] and size in [xs, s, m, l, xl, xxl]."
    exit 1
fi

# Call the function
generate_test_file $1 $2

