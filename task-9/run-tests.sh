#!/bin/sh
root_dir=$(pwd)
find . -name '*_test.go' | while read file; do
	dir=$(dirname -- "$file")
	cd "$dir"
	go test . -v -count=1
	cd "$root_dir"
done
