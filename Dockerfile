FROM golang

RUN mkdir -p /go/src/basic_web_server 
ADD main.go /go/src/basic_web_server
RUN go install basic_web_server
ENTRYPOINT /go/bin/basic_web_server
 
EXPOSE 8080
