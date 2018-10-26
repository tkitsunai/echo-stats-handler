all: vet dep test

vet:
	go vet ./...

dep:
	dep ensure
	dep status

test:
	go test ./...