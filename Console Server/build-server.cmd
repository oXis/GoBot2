go get github.com/gorilla/securecookie
go build -o Server.exe Server.go Crypto.go FileWork.go Session.go
REM to run: Server.exe root root