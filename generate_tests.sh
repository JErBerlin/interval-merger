#!/usr/bin/env bash
# File: generate_tests.sh

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
            num_intervals=10
            ;;
        s)
            num_intervals=50
            ;;
        m)
            num_intervals=200
            ;;
        l)
            num_intervals=5000
            ;;
        xl)
            num_intervals=50000
            ;;
        *)
            echo "Invalid size!"
            exit 1
    esac

    a=1
    b=$((a + max_range - 1))
    
    test_file_path="./testdata/$range.$size.txt"
    expected_file_path="./testdata/expected/$range.$size.txt"

    > $test_file_path
    > $expected_file_path

    for (( i=0; i<$num_intervals; i++ )); do
        point=$((RANDOM % (b - a + 1) + a))
        echo "[$a,$point] [$point,$b] " >> $test_file_path
    done

    echo "[$a,$b]" > $expected_file_path
}

# Call the function
generate_test_file $1 $2

