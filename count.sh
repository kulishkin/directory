#!/usr/bin/bash

dirs=("dir1" "dir2" "dir3")
sum=0
for dir in "${dirs[@]}"
do
  sum="$(($sum+"$(find "$dir" -name 'count' -exec awk '{sum += $1} END {printf "%.0f",sum}' {} +)"))"
done
echo "$sum"
