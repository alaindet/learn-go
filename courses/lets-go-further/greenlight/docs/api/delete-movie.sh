API_URL=http://localhost:4000/v1.0

(
curl \
--request DELETE \
--url $API_URL/movies/7
) | json_pp
