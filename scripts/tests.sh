while true ; do
    if docker exec testdb pg_isready -U postgres ; then
        go test -v  ./src/...
        break
    fi
done