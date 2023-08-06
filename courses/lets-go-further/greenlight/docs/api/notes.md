# Notes

## How to display both response headers and formatted JSON body in cURL
- http://blog.aaronholmes.net/displaying-response-headers-and-pretty-json-with-curl/
- Sample code
  ```console

  # Long form
  curl \
  --silent \
  --dump-header /dev/stderr \
  --url https://jsonplaceholder.typicode.com/todos/1 \
  | jq

  # Short form
  curl -s -D /dev/stderr https://jsonplaceholder.typicode.com/todos/1 | jq
  ```
- What this does is
  1. Remove noisy stats text with `--silent`
  2. Send the response headers to STDERR with `--dump-header /dev/stderr`
  3. Headers will still be shown since STDERR defaults to the terminal
  4. By putting the headers aside, the output of `curl` will simply be the response body
  5. You can then pipe the response body output to `jq` which is a JSON formatter (in case the response body is in JSON, as usual)

## How to declare variables with default values in Bash
```console
GL_API="${GL_API:-http://localhost:4000/v1.0}"
```

- This means "declare `GL_API` variable. If another `GL_API` exists use it, otherwise, use the default value `http://localhost:4000/v1.0`" (mind the `:-`)

## How to create a new test case
```console
export TEST_CASE_FILENAME=my-test-case.sh && ./_new.sh
```

## How to change local test case variables
- With default variables
  ```console
  ./update-movie.sh
  ```
- Override variable `GL_MOVIE_ID`
  ```console
  export GL_MOVIE_ID=5 && ./update-movie.sh
  ```

## How to run two requests sequentially
```console
./healthcheck.sh && ./healthcheck.sh
```

## How to run two requests in parallel
```console
./healthcheck.sh & ./healthcheck.sh
```
