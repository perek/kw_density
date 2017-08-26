.PHONY: build doc fmt lint run test vendor_clean vendor_get vendor_update vet

default: build

build: vet
	go build -v -o ./bin/kwd ./

doc:
	godoc -http=:6060 -index

run:
	./bin/kwd

dev: build
	./bin/kwd

vet:
	go vet ./