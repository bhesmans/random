Small test of the nolibc header from the tools of the linux kernel.

Usage:

```
$ make
gcc -fno-asynchronous-unwind-tables -fno-ident -s -Os  -static -lgcc -o hello_libc hello_libc.c
curl -O https://raw.githubusercontent.com/torvalds/linux/master/tools/include/nolibc/nolibc.h
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 86436  100 86436    0     0   240k      0 --:--:-- --:--:-- --:--:--  240k
gcc -fno-asynchronous-unwind-tables -fno-ident -s -Os -nostdlib  -static -include nolibc.h -lgcc -o hello_nolibc hello_nolibc.c
find . -perm -111 -type f -printf "%p %s B\n"
./hello_nolibc 792 B
./hello_libc 774392 B
diff -q <(./hello_nolibc) <(./hello_libc)
```
