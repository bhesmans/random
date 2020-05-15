Small test about TCPRepairWindow

Note: must have the right to call this sockopt, this is why the Makefile uses
`sudo`

Usage:
```
$ make test
sudo GOPATH=${GOPATH} $(which go) test . -v
=== RUN   TestTCPRepair
2020/05/15 11:54:30 TCP repair 1
2020/05/15 11:54:30 &{Snd_wl1:3833581097 Snd_wnd:65536 Max_window:65536 Rcv_wnd:65408 Rcv_wup:3833581198}
2020/05/15 11:54:30 TCP repair 0
2020/05/15 11:54:30 Resp: Hello !

--- PASS: TestTCPRepair (2.00s)
PASS
ok      _/home/bhesmans/git/random/sockopt      2.006s
```
