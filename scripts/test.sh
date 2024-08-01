#!/bin/sh

base_dir=$(pwd)
for p in `find . -name go.mod | xargs dirname`; do
  cd $p;
  go test -v ./...;
  cd $base_dir;
done
