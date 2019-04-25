#!/bin/bash

set -x

function haha() {
    echo "this is a function..."

    for arg in $@;do
        echo $arg
    done


}

haha 1 2 3 4 5 6
