export GOPATH=$PWD
go get github.com/gin-gonic/gin
go get github.com/smartwalle/pongo2gin
cd src/server
go run main.go