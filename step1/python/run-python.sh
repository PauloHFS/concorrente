#!/bin/bash
echo "" > result

for file in ../../dataset/*
do
    python3 word_count.py $file >> result &
done
wait

awk '{s+=$1} END {print s}' result
