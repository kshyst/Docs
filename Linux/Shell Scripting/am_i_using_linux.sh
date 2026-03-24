#!/bin/bash

kernel=`uname -s`

if [ $kernel = "Linux" ]
then
    echo "YES! You are using linux"
else
    echo "NO! You are NOT using linux"
fi

exit 123
