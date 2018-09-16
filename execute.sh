go get ./...
go build
go test -cover ./...
go run main.go
# export CGO_ENABLED=1
# export GOOS=linux
# go build -o kal-shopping-linux-amd64
# export GOOS=darwin
# export CGO_ENABLED=0