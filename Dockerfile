FROM golang:latest
RUN go version
WORKDIR /go/src/app
ADD . /go/src/app/.
WORKDIR /go/src/app
RUN go get -u
RUN go build -o main .
CMD ["go","run","main.go"]
EXPOSE 8080
# ADD . /go/src/github.com/karuppiah/shopping
# RUN go install github.com/karuppaiah/shopping
# ENTRYPOINT /go/bin/shopping



