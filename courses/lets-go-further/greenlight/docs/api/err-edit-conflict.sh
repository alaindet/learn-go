# Try to edit the same movie 8 times as fast as you can in order to simulate
# concurrency and hit an editing conflict error
xargs -I % -P8 curl -X PATCH -d '{"year": 2006}' "localhost:4000/v1.0/movies/2" < <(printf '%s\n' {1..8})
