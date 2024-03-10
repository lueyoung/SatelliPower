include Makefile.inc

all:
	echo hi

cmain: 
	#gcc -o $@ main.c -I./ ./lib/libsatellite.so
	gcc -o $@ main.c -I./ -L./lib -lm -lsatellite

clean:
	rm -f main
	rm -f cmain
	rm -f app
	rm -f lib/*
	rm -f *.o
	rm -f *.ar

run: build
	./main

compile: ./lib/libsatellite.a
./lib/libsatellite.a: satellite_code.o
	mkdir -p lib
	ar rcs $@ $< 
satellite_code.o: satellite_code.c
	gcc -o $@ -c $< $(CFLAGS)

build: main

main: clean compile
	#CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./$@ $@.go
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./$@ *.go

test:
	docker build -t test-go .
	docker run --rm -it test-go 


.PHONY:
	all clean run compile lib build
