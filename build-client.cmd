go get github.com/gonutz/w32 REM fixed procRedrawWindow for 64 bit
go get github.com/NebulousLabs/go-upnp
go get golang.org/x/sys/windows/registry
go get github.com/StackExchange/wmi
go get github.com/atotto/clipboard
go build -o Client.exe -ldflags "-H windowsgui" GoBot.go 