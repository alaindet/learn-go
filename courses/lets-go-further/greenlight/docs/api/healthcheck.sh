GL_API="${GL_API:-http://localhost:4000/v1.0}"

curl \
--request GET \
--url $GL_API/healthcheck \
--dump-header /dev/stderr \
--silent \
| jq
