xargs -I % -P8 curl -X PATCH -d '{"year": 2006}' "localhost:4000/v1.0/movies/2" < <(printf '%s\n' {1..8})
