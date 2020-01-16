#!/bin/bash
x=0
while true 
do
	mod=$(( $x % 100 ))
	if [ $mod -eq 0 ]
	then
		echo "Done $x times"
	fi
	x=$(( $x + 1 ))
	eval './build/rubik -r 10000 -re n -a 4'
	if [ $? -ne 0 ]
	then
		exit 1
	fi
done
