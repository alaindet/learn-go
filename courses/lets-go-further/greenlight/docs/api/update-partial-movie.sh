API_URL=http://localhost:4000/v1.0

(
curl \
--request PATCH \
--url $API_URL/movies/2 \
--header 'Content-Type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
  "title": "The Matrix"
}
EOF
) | json_pp
