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

i=1
sp="/-\|"
echo -n ' '
progress="..."
echo -ne "\n"
while true ; do
    echo -ne "\r=== Waiting for test database to connect  \b${sp:i++%${#sp}:1} ${progress}"
    myfunc
    if  [ $? -eq "1" ]; then
        echo -ne "\n=== Test database is ready\n\n"
        go test -v  ./src/...
        break
    fi
    progress="${progress}."
done


