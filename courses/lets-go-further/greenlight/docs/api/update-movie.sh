API_URL=http://localhost:4000/v1.0

(
curl \
--request PUT \
--url $API_URL/movies/1 \
--header 'Content-Type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
  "title": "The Matrix",
  "year": 1999,
  "runtime": "136 mins",
  "genres": ["sci-fi", "action"]
}
EOF
) | json_pp
