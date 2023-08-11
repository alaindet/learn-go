GL_API="${GL_API:-http://localhost:4000/v1}"

for i in {1..10}; do curl $GL_API/healthcheck; done | jq
