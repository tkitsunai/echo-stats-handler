all: vet glide test

vet:
	go vet ./...

glide:
	glide install

test:
	go test ./...