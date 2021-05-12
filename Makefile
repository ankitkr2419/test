test:
	go test -v ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	sudo ./build-binary.sh

zip:
	./build-zip.sh
	echo "zip created successfully"

build-and-zip: build zip