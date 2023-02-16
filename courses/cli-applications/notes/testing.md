# Testing notes

## Fixture
- A test **fixture** is just an environment that replicates a specific condition in which the test can be run meaningfully

## Golden file
- A **golden file** is a reference file containing expected outcome to be compared with
- It must be versioned
- It's used when testing functions that generate files

## `testdata`
- Folders named `testdata` are excluded by the Go builder by default
- `testdata` contains test files, usually golden and configuration files
