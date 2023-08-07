GL_API=http://localhost:4000/v1.0

curl \
--request GET \
--url "$GL_API/movies?page=1&page_size=10&genres=gen1,gen2&title=Some%20Title&sort=year" \
--dump-header /dev/stderr \
--silent \
| jq
