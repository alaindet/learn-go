curl \
  -X POST \
  -d '{ "title": "Catch Me If You Can", "year": 2002, "runtime": "141 mins", "genres": ["biography", "comedy", "drama"] }' \
  http://localhost:4000/v1.0/movies \
  | json_pp
