test:
	go test -v ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	sudo utils/build-yarn.sh
	echo "building go code"
	GOARCH=amd64 GOOS=linux
	go build -v -ldflags=" \
	-X 'main.Version=v1.2.1' \
	-X 'main.User=$(shell id -u -n)' \
	-X 'main.BuiltOn=$(shell date)' \
	-X 'main.CommitID=$(shell git rev-list -1 HEAD)' \
	-X 'main.Branch=$(shell git rev-parse --abbrev-ref HEAD)' \
	-X 'main.Machine=$(shell hostname)'"
	echo "binary created"

zip:
	utils/build-zip.sh
	echo "zip created successfully"

build-and-zip: build zip
