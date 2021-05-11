test:
	go test -v ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

build-and-zip:
	sudo ./build-binary.sh
	zip mylab.zip migrations conf cpagent installation.sh
	echo "zip created"

zip:
	zip mylab.zip migrations conf cpagent installation.sh
	echo "zip created"