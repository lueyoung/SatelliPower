all:
	@echo Hi

sender:
	./src/$@/$@.py

recver:
	./src/$@/$@.py

midware: main
	./src/$@/$<

main:
	cd ./src/midware; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./$@ *.go

clean:
	lsof -i :8080 | egrep '[0-9]+' | awk -F ' ' '{print $2}' | xargs -I {} kill -9 {}
	lsof -i :8081 | egrep '[0-9]+' | awk -F ' ' '{print $2}' | xargs -I {} kill -9 {}

.PHONY: all run up clean down build sender recver midware

