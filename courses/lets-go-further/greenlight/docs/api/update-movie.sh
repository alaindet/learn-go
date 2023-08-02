curl \
  -X PUT \
  -d '{ "title": "The Matrix", "year": 1999, "runtime": "136 mins", "genres": ["sci-fi", "action"] }' \
  http://localhost:4000/v1.0/movies/1 \
  | json_pp
