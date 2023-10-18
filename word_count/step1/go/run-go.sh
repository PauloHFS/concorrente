#!/bin/bash
echo "" > result

for file in ../../dataset/*
do
    go run word_count.go $file >> result &
done
wait

awk '{s+=$1} END {print s}' result
