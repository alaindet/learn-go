GL_API="${GL_API:-http://localhost:4000/v1}"

curl \
--request POST \
--url $GL_API/movies \
--header 'Accept: application/json' \
--header 'Content-Type: application/json; charset=utf-8' \
--dump-header /dev/stderr \
--silent \
--data '{
  "title": "Donnie Darko",
  "year": 2001,
  "runtime": "113 mins",
  "genres": ["some-genre"]
}' \
| jq
