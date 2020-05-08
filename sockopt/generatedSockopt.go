package sockopt

// #include <unistd.h>
import "C"

import "syscall"
import "unsafe"

type _Socklen C.socklen_t

// Errno values.
var (
	errEAGAIN error = syscall.EAGAIN
	errEINVAL error = syscall.EINVAL
	errENOENT error = syscall.ENOENT
)

// Errors
const (
	E2BIG       = syscall.Errno(0x7)
	EACCES      = syscall.Errno(0xd)
	EAGAIN      = syscall.Errno(0xb)
	EBADF       = syscall.Errno(0x9)
	EBUSY       = syscall.Errno(0x10)
	ECHILD      = syscall.Errno(0xa)
	EDOM        = syscall.Errno(0x21)
	EEXIST      = syscall.Errno(0x11)
	EFAULT      = syscall.Errno(0xe)
	EFBIG       = syscall.Errno(0x1b)
	EINTR       = syscall.Errno(0x4)
	EINVAL      = syscall.Errno(0x16)
	EIO         = syscall.Errno(0x5)
	EISDIR      = syscall.Errno(0x15)
	EMFILE      = syscall.Errno(0x18)
	EMLINK      = syscall.Errno(0x1f)
	ENFILE      = syscall.Errno(0x17)
	ENODEV      = syscall.Errno(0x13)
	ENOENT      = syscall.Errno(0x2)
	ENOEXEC     = syscall.Errno(0x8)
	ENOMEM      = syscall.Errno(0xc)
	ENOSPC      = syscall.Errno(0x1c)
	ENOTBLK     = syscall.Errno(0xf)
	ENOTDIR     = syscall.Errno(0x14)
	ENOTTY      = syscall.Errno(0x19)
	ENXIO       = syscall.Errno(0x6)
	EPERM       = syscall.Errno(0x1)
	EPIPE       = syscall.Errno(0x20)
	ERANGE      = syscall.Errno(0x22)
	EROFS       = syscall.Errno(0x1e)
	ESPIPE      = syscall.Errno(0x1d)
	ESRCH       = syscall.Errno(0x3)
	ETXTBSY     = syscall.Errno(0x1a)
	EWOULDBLOCK = syscall.Errno(0xb)
	EXDEV       = syscall.Errno(0x12)
)

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT
// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case EAGAIN:
		return errEAGAIN
	case EINVAL:
		return errEINVAL
	case ENOENT:
		return errENOENT
	}
	return e
}

// ErrnoName returns the error name for error number e.