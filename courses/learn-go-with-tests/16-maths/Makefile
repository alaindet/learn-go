.DEFAULT_GOAL := build

test:
	go test learn_go_with_tests/maths/clockface

build:
	go build -o "./dist/clockface"
	./dist/clockface > ./dist/clock.svg
