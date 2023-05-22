#!/bin/bash

for i in $(seq 0 113);
do
    s=$(cat chapters.json | jq ".chapters[$i].pages[0]")
    e=$(cat chapters.json | jq ".chapters[$i].pages[1]")
    echo "$i $s $e"
done
