.PHONY: all update-modules

all:
	env GO111MODULE=on go build -v ./...

test:
	gocheck -g -c -e _example -e sqlite3_test

update-modules:
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
