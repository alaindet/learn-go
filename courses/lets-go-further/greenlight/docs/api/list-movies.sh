GL_API=http://localhost:4000/v1.0

# curl \
# --request GET \
# --url "$GL_API/movies" \
# --dump-header /dev/stderr \
# --silent \
# | jq

curl \
--request GET \
--url "$GL_API/movies?page=2&page_size=5&genres=gen1,gen2&title=Some%20Title&sort=-year" \
--dump-header /dev/stderr \
--silent
