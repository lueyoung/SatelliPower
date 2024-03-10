all:
	echo hi

main: 
	gcc -o $@ $@.c -I./ ./lib/libsatellite.so

clean:
	rm -f main
	rm -f app
	rm -f lib/*

run: compile
	make set_env && go run main.go

compile:
	gcc -o ./lib/libsatellite.so satellite_code.c -fPIC -shared -I./ -lm

lib_path = $(shell pwd)
set_env:
	export LD_LIBRARY_PATH=$(LD_LIBRARY_PATH):$(lib_path)/lib

exe:
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./app *.go

format:
	-find ./ -type d | grep -v proto | xargs -I {} find {} -type f | egrep "\.h$$|\.cpp$$" | xargs -I {} indent -kr -i4 {}

