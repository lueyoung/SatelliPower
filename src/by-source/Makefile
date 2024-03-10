all:
	echo hi

clean:
	rm -f main
	rm -f cmain
	rm -f app
	rm -f lib/*
	rm -f *.o
	rm -f *.ar

run: build
	./main

build: main

main: clean
	#CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./$@ $@.go
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./$@ *.go

test:
	docker build -t test-go .
	docker run --rm -it test-go 


.PHONY:
	all clean run compile lib build
