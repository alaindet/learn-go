GL_API="${GL_API:-http://localhost:4000/v1}"

# Valid
# Running it twice triggers a duplicated email error
curl \
--request POST \
--url $GL_API/users \
--header 'Accept: application/json' \
--header 'Content-Type: application/json; charset=utf-8' \
--dump-header /dev/stderr \
--silent \
--data '{
  "name": "Alice",
  "email": "alice@example.com",
  "password": "alice@example.com"
}' \
| jq

# # Invalid JSON input
# curl \
# --request POST \
# --url $GL_API/users \
# --header 'Accept: application/json' \
# --header 'Content-Type: application/json; charset=utf-8' \
# --dump-header /dev/stderr \
# --silent \
# --data '{
#   "name": "Alice",
#   ""email": "alice@example.com"
#   "password": "alice@example.com",
# }' \
# | jq

# # Invalid data (invalid email)
# curl \
# --request POST \
# --url $GL_API/users \
# --header 'Accept: application/json' \
# --header 'Content-Type: application/json; charset=utf-8' \
# --dump-header /dev/stderr \
# --silent \
# --data '{
#   "name": "Alice",
#   "email": "Alice",
#   "password": "alice@example.com"
# }' \
# | jq
