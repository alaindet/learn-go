GL_API="${GL_API:-http://localhost:4000/v1}"

# # Without filters
# curl \
# --request GET \
# --url "$GL_API/movies" \
# --dump-header /dev/stderr \
# --silent \
# | jq

# # With genres filter
# curl \
# --request GET \
# --url "$GL_API/movies?page=1&page_size=10&genres=comedy&sort=year" \
# --dump-header /dev/stderr \
# --silent \
# | jq

# # Empty query
# curl \
# --request GET \
# --url "$GL_API/movies?page=1&page_size=10&title=NOPE&sort=year" \
# --dump-header /dev/stderr \
# --silent \
# | jq

# # With partial (word-sized) title filter
# curl \
# --request GET \
# --url "$GL_API/movies?title=the" \
# --dump-header /dev/stderr \
# --silent \
# | jq

# With sorting
curl \
--request GET \
--url "$GL_API/movies?sort=-runtime&page=1&page_size=10" \
--dump-header /dev/stderr \
--silent \
| jq
