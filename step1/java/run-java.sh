#!/bin/bash
echo "" > result

for file in ../../dataset/*
do
    gradle runApp --args=$file -q >> result &
done
wait

awk '{s+=$1} END {print s}' result
