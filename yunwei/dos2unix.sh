#!/bin/bash
function my_dos2unix() {
    for i in  `ls $1 `; do
        if [ -d "$1/$i" ]; then
            my_dos2unix "$1/$i"
        else 
            /usr/bin/dos2unix "$1/$i"
            echo "$1/$i"
        fi
    done


}
my_dos2unix $1
