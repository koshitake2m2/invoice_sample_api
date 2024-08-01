#!/bin/sh

base_dir=$(pwd)
for p in `find . -name go.mod | xargs dirname`; do
  cd $p;
  result_p=$(gofmt -l .)
  [ -z "${result_p}" ] || echo $result_p;
  cd $base_dir;
done
