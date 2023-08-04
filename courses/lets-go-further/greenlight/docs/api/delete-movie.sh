GL_API="${GL_API:-http://localhost:4000/v1.0}"

curl \
--request DELETE \
--url $GL_API/movies/1 \
--header 'Accept: application/json' \
--dump-header /dev/stderr \
--silent \
| jq
