all:
	echo hi

cmain: 
	gcc -o $@ main.c -I./ ./lib/libsatellite.so

clean:
	rm -f main
	rm -f cmain
	rm -f app
	rm -f lib/*

run: build
	./main

compile:
	mkdir -p lib
	gcc -o ./lib/libsatellite.so satellite_code.c -fPIC -shared -I./ -lm

build: main

main: compile
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./$@ $@.go

.PHONY:
	all clean run compile lib 
