.PHONY: docker

vendor-update:
	govendor fetch +external +missing

docker:
	docker-compose up -d

docker-down:
	docker-compose down