#!/bin/bash
NITER="${1:-10}"
for i in $(seq -w "$NITER")
do
  echo "$i"
  sleep 1
done