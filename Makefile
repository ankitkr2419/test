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

db-migrate:
	echo "restarting postgresql to avoid connection error"
	sudo service postgresql restart
	echo "dropping cpagentdb"
	echo "you may need to enter database password"
	sudo dropdb -h localhost -p 5432 -U postgres cpagentdb
	echo "creating new cpagentdb database"
	echo "you may need to enter database password"
	sudo -u postgres createdb --owner=postgres cpagentdb
	echo "migrating to cpagentdb"
	./cpagent migrate
	echo "starting simulator to fill up table fields"
	(timeout 3 ./cpagent start --plc simulator; exit 0)
	echo "sleeping for 3 secs"
	sleep 3
	echo "stopping simulator process"
