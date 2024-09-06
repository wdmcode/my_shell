#!/bin/sh

#
# 该脚本用来替换 sql 中的某个值，替换的值来自 key.txt 文件中对的行的值
#

while read line; do
    # echo $line
    i=$((i + 1))

    # key1=$(sed -n "$((i * 2 - 1))p" key.txt)
    key2=$(sed -n "$((i * 2))p" key.txt)

    # echo $key1
    # echo $key2

    echo "$line" | sed "s#\('[^']*','[^']*','\)[^']*\('.*\)#\1${key2}\2#"
done <sql.txt
