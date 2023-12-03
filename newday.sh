#!/bin/bash

# Check if the argument is a valid number between 1 and 25
if [[ $1 =~ ^[0-9]+$ ]] && [ $1 -ge 1 ] && [ $1 -le 25 ]; then
    # Zero-pad the number to two digits
    folder_num=$(printf "%02d" $1)

    # Create directory if it doesn't exist
    dir_name="day$folder_num"
    if [ ! -d "$dir_name" ]; then
        mkdir "$dir_name"
    fi

    # Copy day_test.txt and rename
    if [ ! -f "$dir_name/day${folder_num}_test.go" ]; then
        cp "day_test.txt" "$dir_name/day${folder_num}_test.go"
        sed -i '' "1s/.*/package day$folder_num/" "$dir_name/day${folder_num}_test.go"
    fi

    # Copy day.txt and rename
    if [ ! -f "$dir_name/day${folder_num}.go" ]; then
        cp "day.txt" "$dir_name/day${folder_num}.go"
        # Modify package line
        sed -i '' "1s/.*/package day$folder_num/" "$dir_name/day${folder_num}.go"
        # Modify ReadInput line
        sed -i '' "s/util.ReadInput(\"day.txt\")/util.ReadInput(\"day$folder_num.txt\")/" "$dir_name/day${folder_num}.go"
    fi

    #create input file
    if [ ! -f "input/day${folder_num}.txt" ]; then
        touch "input/day${folder_num}.txt"
    fi
else
    echo "Error: Please provide a number between 1 and 25."
    exit 1
fi
