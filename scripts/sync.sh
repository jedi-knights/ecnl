#!/usr/bin/env sh

mkdir -p output

start=$(date +%s)
go run main.go sync > output/sync.log
end=$(date +%s)
echo "Elapsed Time: $(($end-$start)) seconds"
