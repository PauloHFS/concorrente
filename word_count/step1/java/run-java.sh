#!/bin/bash
cd ./src/main/java

echo "" > result

for file in ../../../../../dataset/*
do
    # agora bora trocar esse comando aqui
    java WordCount $file >> result &
done
wait

awk '{s+=$1} END {print s}' result
