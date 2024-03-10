all:
	echo hi

main: 
	gcc -o $@ $@.c -I./ ./lib/libsatellite.so

clean:
	rm -f main
	rm -f lib/*

run: clean main
	./main

compile:
	gcc -o ./lib/libsatellite.so satellite_code.c -fPIC -shared -I./ -lm
