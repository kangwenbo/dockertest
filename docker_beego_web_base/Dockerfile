FROM golang:1.7.3
RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee
ADD web.go /go/src/web.go
CMD ["go","run","/go/src/web.go"]
