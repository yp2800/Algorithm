#!/usr/bin/env bash

awk '{for(i=1;i<=NF;i++)if(arr[i]){arr[i]=arr[i]" "$i}else{arr[i]=$i}}END{for(key in arr)print arr[key]}' file.txt
