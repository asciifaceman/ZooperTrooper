.PHONY: docker build

vendor-update:
	govendor fetch +external +missing

docker:
	docker-compose up -d

docker-down:
	docker-compose down

build: build-linux build-darwin

build-linux:
	@GOOS=linux GOARCH=amd64 go build

build-darwin:
	@GOOS=darwin GOARCH=amd64 go build