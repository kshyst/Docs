#!/bin/bash

VAR=70

while [ $VAR -gt 55 ]
do
    echo now var is $VAR
    let VAR=VAR-1
done
