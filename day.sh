#!/usr/bin/env bash

year="2023"

default_day="$(date +%d)"
day="${1:-$default_day}"
dir="d$day"

session="$(cat .sessionid)"

echo "Creating directory $dir"
mkdir -p "$dir"
echo "Trying to download input"
curl "https://adventofcode.com/$year/day/$day/input" -H "Cookie: session=$session" -o "$dir/input" || echo "Failed to fetch input"
