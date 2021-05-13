test:
	go test -v ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	sudo utils/build-binary.sh

zip:
	utils/build-zip.sh
	echo "zip created successfully"

build-and-zip: build zip