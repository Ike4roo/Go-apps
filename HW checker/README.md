# Hardware checker for Debian
## Description
Collects info about CPU, RAM, LAN, WLAN, USB devices, GPU attached and stores in `*.txt` file and shows on a screen using dialog

## Requirements
- Golang installed
- Dialog installed (easy as `apt install dialog`)
- Some skills and Internet

Depends on 3D party libs (look into *.go file)

## Make it
Before compile, be sure Golang and git are installed on a machine where you build this code.

Import 3d-party libraies:
```bash
go mod init hw_info
go get github.com/shirou/gopsutil
go get github.com/guptarohit/asciigraph
```

Build:
`go build -o hw_info`

Execute:
`./hw_info`
