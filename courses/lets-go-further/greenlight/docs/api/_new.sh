TEST_CASE_FILENAME="${TEST_CASE_FILENAME:-my-test-case.sh}"

echo -e "GL_API="${GL_API:-http://localhost:4000/v1.0}"

curl \\
--request GET \\
--url \$GL_API/healthcheck \\
--dump-header /dev/stderr \\
--silent \\
| jq" \
> $TEST_CASE_FILENAME && \
chmod +x $TEST_CASE_FILENAME
