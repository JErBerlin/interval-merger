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

    # Calculate the average diameter of intervals
    avg_diameter=$((max_range / num_points))
    variability=$((avg_diameter)) # up to 100% of the average diameter

    # Write the number of intervals first
    echo $((2 + ($num_points - 2) * 2)) > $test_file_path
    echo "1" > $expected_file_path

    # Adjusting calculation for point1 and point2 using average diameter with variability
    dia_variability1=$((RANDOM % (2 * variability + 1) - variability))
    dia_adjusted1=$((avg_diameter + dia_variability1))
    
    point1=$((a + dia_adjusted1))
    
    dia_variability2=$((RANDOM % (2 * variability + 1) - variability))
    dia_adjusted2=$((avg_diameter + dia_variability2))
    
    point2=$((b - dia_adjusted2))

    echo "[$a,$point1]" >> $test_file_path
    echo "[$point2,$b]" >> $test_file_path

    for (( i=0; i<($num_points-2)*2; i++ )); do
        dia_variability=$((RANDOM % (2 * variability + 1) - variability)) 
        dia_adjusted=$((avg_diameter + dia_variability))
        
        a_i=$((RANDOM % (b - dia_adjusted - a + 1) + a))
        b_i=$((a_i + dia_adjusted))

        # Append to a new line for the test file
        echo "[$a_i,$b_i]" >> $test_file_path
    done

    echo "[$a,$b]" >> $expected_file_path
}

# Call the function
generate_test_file $1 $2

