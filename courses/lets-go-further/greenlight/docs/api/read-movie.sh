GL_API="${GL_API:-http://localhost:4000/v1}"
GL_MOVIE_ID="${GL_MOVIE_ID:-1}"

curl \
--request GET \
--url $GL_API/movies/$GL_MOVIE_ID \
--header 'Accept: application/json' \
--dump-header /dev/stderr \
--silent \
| jq
