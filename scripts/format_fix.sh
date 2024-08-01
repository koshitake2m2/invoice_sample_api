#!/bin/sh

base_dir=$(pwd)
for p in `find . -name go.mod | xargs dirname`; do
  cd $p;
  go fmt
  cd $base_dir;
done
