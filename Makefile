all:
	go build -o calcuCo
test:
	go test -v ./...
benchmark:
	go test -bench . ./...