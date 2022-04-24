#!/bin/bash

LANG_ENG="accepting connections"
LANG_PT_BR="aceitando conexões"

myfunc()
{
    langs=( $LANG_ENG $LANG_PT_BR)
    v=$(docker exec testdb pg_isready -U postgres)
    for l in "${langs[@]}";do
        if [[ $v == *"$l"* ]]; then
            return 1
        else
            return 2
        fi
    done
}

SECONDS_LIMIT=10
i=1
sp="/-\|"
echo -n ' '
progress="..."
echo -ne "\n"
START=$(date +%Y%m%d%H%M%S)
while true ; do
    echo -ne "\r=== TEST DATABASE   waiting for connection  \b${sp:i++%${#sp}:1} ${progress}"
    myfunc
    if  [ $? -eq "1" ]; then
        echo -ne "\n=== TEST DATABASE   connection is ready. \n\n"
        go test -v  ./src/...
        break
    fi

    CURRENT=$(date +%Y%m%d%H%M%S)
    limit=$(($CURRENT - $START))
    if [ $limit -gt $SECONDS_LIMIT ];then
        echo -ne "\n=== TEST DATABASE   timed out when trying, limit ${SECONDS_LIMIT} seconds. \n\n"
        break
    fi

    progress="${progress}."

done


