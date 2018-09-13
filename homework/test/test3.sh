#!/bin/bash


[ $(uname -r | cut -c1) -eq 2 -a $(uname -r | cut -c3) -ge 6 ] && echo $(uname -r)
