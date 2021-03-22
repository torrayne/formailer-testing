#!/bin/bash
dir=$(dirname $1)
dir=${dir##*/}
go build -o "functions/$dir" $1
