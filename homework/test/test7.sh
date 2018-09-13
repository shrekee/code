#!/bin/bash

[ ! -d /rhome ] && $(mkdir /rhome)

for i in stu{1..5};do
	useradd -d /rhome $i > /dev/null
done

