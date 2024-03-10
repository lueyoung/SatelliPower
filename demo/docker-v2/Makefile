all:
	@echo Hi

run: up

up: clean
	docker-compose $@ -d

clean: down

down:
	docker-compose $@

build:
	docker-compose $@

sender:
	docker logs -f $@
recver:
	docker logs -f $@
midware:
	docker logs -f $@

.PHONY: all run up clean down build sender recver midware

