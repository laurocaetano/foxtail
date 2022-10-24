# run tests
test:
	go mod tidy
	go fmt ./...
	go test ./...

# build
build: | test
	go build -o ./pkg/foxtail

# compile
compile:
	go fmt ./...
	go build ./...
