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
	-X 'mylab/cpagent/service.Version=v1.4.178' \
	-X 'mylab/cpagent/service.User=$(shell id -u -n)' \
	-X 'mylab/cpagent/service.BuiltOn=$(shell date)' \
	-X 'mylab/cpagent/service.CommitID=$(shell git rev-list -1 HEAD)' \
	-X 'mylab/cpagent/service.Branch=$(shell git rev-parse --abbrev-ref HEAD)' \
	-X 'mylab/cpagent/service.Machine=$(shell hostname)'"
	echo "binary created"

zip:
	utils/build-zip.sh
	echo "zip created successfully"

baz: build-and-zip

build-and-zip: build zip

migrate:
	utils/db-migrate.sh

upgrade-tec-api:
	utils/upgrade-tec-api.sh
