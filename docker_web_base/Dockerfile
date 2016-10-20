FROM golang:1.7.3

ADD main.go /go/src/main.go
ADD run.sh /run.sh
RUN chmod +x /run.sh
CMD ["/bin/bash","/run.sh"]
#EXPOSE 8080
