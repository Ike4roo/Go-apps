# Screensaver for desktop (Debian xfce) that checks application (in tjis case is for Home Assistant) while shows any animation

Is used by me to show an animation `*.gif` in a fullscreen (hiding all OS interface) to not to let users interract with desktop (without any safety or blocking keyboard or mouse) while it waits for Home Assitant ready on a dedicated address

This small program just checks availability of application using HTTP GET method, so could be used with any application that has it

## Prerequisites
- Debian (or Debian-like) OS
- XFCE desktop (or any)
- Desktop browser
- Network link up and DNS resolver
- Go compiler and Git client packet installed (just for build)
- libX11, libGL, libasound (for github.com/hajimehoshi/ebiten/v2)

**Depends on github.com/hajimehoshi/ebiten/v2**

## How-to

Open `*.go` file. Find and change `openBrowser("http://192.168.11.12:8123")` with IP or FQDN and port you like, this program will to check

Put any animation saved as `animation.gif` file near to `*.go` file

Build it and execute
```bash
sudo apt update
sudo apt install -y libx11-dev libxext-dev libxxf86vm-dev
go get github.com/hajimehoshi/ebiten/v2
go build -o home_assistant_checker main.go
```

Check if your Debian OS has all dependencies installed for binary file you've just compiled:
```bash
ldd <binary_file_name>
```