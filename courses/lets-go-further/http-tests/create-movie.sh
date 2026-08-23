curl -i -X POST http://localhost:4000/v1/movies -d '{
  "title":"Moana",
  "year":2016,
  "runtime":107,
  "genres":["animation","adventure"]
}'