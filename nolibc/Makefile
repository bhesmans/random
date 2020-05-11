.PHONY: printSize test clean

all: hello_libc hello_nolibc printSize test

#dl the nolibc header from kernel's tool src
#gcc command extracted from the comments of this file
nolibc.h:
	curl -O https://raw.githubusercontent.com/torvalds/linux/master/tools/include/nolibc/nolibc.h

hello_nolibc: hello_nolibc.c nolibc.h
	gcc -fno-asynchronous-unwind-tables -fno-ident -s -Os -nostdlib  -static -include nolibc.h -lgcc -o $@ $<

hello_libc: hello_libc.c
	gcc -fno-asynchronous-unwind-tables -fno-ident -s -Os  -static -lgcc -o $@ $<

printSize: hello_nolibc hello_libc
	find . -perm -111 -type f -printf "%p %s B\n"

test: SHELL := /bin/bash
test: hello_nolibc hello_libc
	diff -q <(./hello_nolibc) <(./hello_libc)

clean:
	rm -rf hello_nolibc hello_libc nolibc.h
