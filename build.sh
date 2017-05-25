export GOPATH=$PWD
go get github.com/gin-gonic/gin
go get github.com/smartwalle/pongo2gin
go get github.com/golang/net/context
cd src/server
go run main.go
