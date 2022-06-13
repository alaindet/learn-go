# Testing

- Any file with a name like `*_test.go` is excluded by the builder, does not get imported and cannot be used
- For this reason, since we're providing stubs/mocks here and testing is involved, we named the files `test_*.go` instead

- **TODO**: Should there be a separate testing package named `data_test`?
