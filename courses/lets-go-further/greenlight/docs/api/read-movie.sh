API_URL=http://localhost:4000/v1.0

(
curl \
--request GET \
--url $API_URL/movies/1
) | json_pp
