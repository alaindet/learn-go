API_URL=http://localhost:4000/v1.0

(
curl \
--request POST \
--url $API_URL/movies \
--header 'Content-Type: application/json; charset=utf-8' \
--data-binary @- << EOF
{
  "title": "Catch Me If You Can",
  "year": 2002,
  "runtime": "141 mins",
  "genres": ["biography", "comedy", "drama"]
}
EOF
) | json_pp
