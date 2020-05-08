package sockopt

// #include <linux/tcp.h>
import "C"

import "unsafe"

type TCPRepairWindow struct {
	Snd_wl1    uint32
	Snd_wnd    uint32
	Max_window uint32
	Rcv_wnd    uint32
	Rcv_wup    uint32
}

var SizeofTCPRepairWindow = C.sizeof_struct_tcp_repair_window

func GetsockoptTCPRepairWindow(fd, level, opt int) (*TCPRepairWindow, error) {
	var value TCPRepairWindow
	vallen := _Socklen(SizeofTCPRepairWindow)
	err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
	return &value, err
}
