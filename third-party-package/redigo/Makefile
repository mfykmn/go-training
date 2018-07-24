## Docker
CONTAINER_PREFIX:=redigo_

.PHONY: dstart dstop dstatus dlogin dclean dlog
dstart:
	@echo "docker start"
	@docker-compose up -d

dstop:
	@echo "docker stop"
	@docker-compose stop

dstatus:
	@echo "docker status"
	@docker ps --filter name=$(CONTAINER_PREFIX)

dlogin:
	@echo "docker login"
	@docker exec -it $(shell docker ps --all --format "{{.Names}}" | peco) /bin/bash

dclean:
	@echo "docker clean"
	@docker ps --all --filter name=$(CONTAINER_PREFIX) --quiet | xargs docker rm --force
	@rm -rf ./minio-data

dlog:
	@echo "docker log"
	@docker-compose logs -f $(shell docker ps --all --format "{{.Names}}" | peco | cut -d"_" -f2)
