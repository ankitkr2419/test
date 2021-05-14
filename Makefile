test:
	go test -v ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	sudo utils/build-binary.sh
	echo "building go code"
	GOARCH=amd64 GOOS=linux
	go build -v -ldflags=" \
	-X 'main.Version=v1.2.1' \
	-X 'main.User=$(shell id -u -n)' \
	-X 'main.Built=$(shell date)' \
	-X 'main.CommitID=$(shell git rev-list -1 HEAD)'"
	echo "binary created"

zip:
	utils/build-zip.sh
	echo "zip created successfully"

build-and-zip: build zip
