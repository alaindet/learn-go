GL_API="${GL_API:-http://localhost:4000/v1.0}"
GL_MOVIE_ID="${GL_MOVIE_ID:-1}"

# # All info
# curl \
# --request PATCH \
# --url $GL_API/movies/$GL_MOVIE_ID \
# --header 'Accept: application/json' \
# --header 'Content-Type: application/json; charset=utf-8' \
# --dump-header /dev/stderr \
# --silent \
# --data '{
#   "title": "The Matrix",
#   "year": 1999,
#   "runtime": "136 mins",
#   "genres": ["sci-fi", "action"]
# }' \
# | jq

# Partial info
curl \
--request PATCH \
--url $GL_API/movies/$GL_MOVIE_ID \
--header 'Accept: application/json' \
--header 'Content-Type: application/json; charset=utf-8' \
--dump-header /dev/stderr \
--silent \
--data '{
  "title": "Some New Movie Name"
}' \
| jq

# # Invalid (cannot pass empty fields)
# curl \
# --request PATCH \
# --url $GL_API/movies/$GL_MOVIE_ID \
# --header 'Accept: application/json' \
# --header 'Content-Type: application/json; charset=utf-8' \
# --dump-header /dev/stderr \
# --silent \
# --data '{
#   "title": "",
#   "year": 2000
# }' \
# | jq
