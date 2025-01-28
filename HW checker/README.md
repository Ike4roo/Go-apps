# Hardware checker for Debian
## Description
Collects info about CPU, RAM, LAN, WLAN, USB devices, GPU attached and stores in `*.txt` file and shows on a screen

## Make it
Before compile, be sure Golang and git are installed on a machine where you build this code.

Import 3d-party libraies:
```bash
go get github.com/shirou/gopsutil
go get github.com/guptarohit/asciigraph
```

Build:
`go build -o hw_info`

Execute:
`./hw_info`