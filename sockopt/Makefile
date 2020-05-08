.PHONY: generatedSockopt.go

GWT = ${GOPATH}/src/golang.org/x/sys
GD = $(GWT)/.git

define get_fun
	GIT_WORK_TREE=$(GWT) GIT_DIR=$(GD) git grep -W "func $(1)" -- $(2) | sed 's#$(2).##'
endef

generatedSockopt.go:
	echo "package sockopt\n" > $@
	echo '// #include <unistd.h>\nimport "C"\n' >> $@
	echo 'import "syscall"' >> $@
	echo 'import "unsafe"\n' >> $@
	echo 'type _Socklen C.socklen_t\n' >> $@
	GIT_WORK_TREE=$(GWT) GIT_DIR=$(GD) git grep -C2 "error =" -- unix/syscall_unix.go | sed 's#unix/syscall_unix.go.##' >> $@
	GIT_WORK_TREE=$(GWT) GIT_DIR=$(GD) git grep -C2 "syscall.Errno" -- unix/zerrors_linux.go | sed 's#unix/zerrors_linux.go.##' >> $@
	$(call get_fun,getsockopt,unix/zsyscall_linux_amd64.go) >> $@
	$(call get_fun,errnoErr,unix/syscall_unix.go) >> $@
	sed -i 's/Syscall6/syscall.Syscall6/' $@
	sed -i 's/SYS_GETSOCKOPT/syscall.SYS_GETSOCKOPT/' $@
	go build ./$@

