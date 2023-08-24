#!/bin/bash

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
            num_points=5
            ;;
        s)
            num_points=25
            ;;
        m)
            num_points=100
            ;;
        l)
            num_points=2500
            ;;
        xl)
            num_points=25000
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

    for (( i=0; i<$num_points; i++ )); do
        point=$((RANDOM % (b - a + 1) + a))
        
        a_i=$((RANDOM % (point - a + 1) + a))
        b_i=$((RANDOM % (b - point + 1) + point))

        a_i_next=$((RANDOM % (point - a + 1) + a))
        b_i_next=$((RANDOM % (b - point + 1) + point))
        
        # Append to the same line for the test file
        printf "[%d,%d] [%d,%d] " $a_i $b_i $a_i_next $b_i_next >> $test_file_path
    done

    # Add newline to the end of the test file
    echo "" >> $test_file_path

    echo "[$a,$b]" > $expected_file_path
}

# Call the function
generate_test_file $1 $2

