Toy project that exposes CPU load with a simple HTTP server. CPU loads are sent
through websocket to update the animated gauges. CPU loads are extracted from
`/proc/stat`

Uses Gorilla for ws : https://github.com/gorilla/websocket
Uses this lib for the gauges : https://www.cssscript.com/animated-svg-gauge/

Usage:

```
go build ./main.go
./main
# Browse localhost:8080
```

![Demo](demo.png?raw=true "Demo")
