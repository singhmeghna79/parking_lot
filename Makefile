.PHONY: clean test build
.DEFAULT_GOAL := build

build:
	cd cmd/parking_lot && go build -o ../../bin/parking_lot

clean:
	rm -f ./bin/parking_lot

test:
	go test ./... -cover
