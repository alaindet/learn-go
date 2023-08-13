GL_API="${GL_API:-http://localhost:4000/v1}"

curl \
--request POST \
--url $GL_API/users \
--header 'Accept: application/json' \
--header 'Content-Type: application/json; charset=utf-8' \
--dump-header /dev/stderr \
--silent \
--data '{
  "email": "user@example.com"
}' \
| jq
