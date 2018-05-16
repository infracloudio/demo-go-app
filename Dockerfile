ARG GO_VERSION=1.9.2
FROM golang:${GO_VERSION}

ENV GOPATH /usr
ENV APP	   ${GOPATH}/src/github.com/infracloudio/demo-go-app/

ADD main.go ${APP}
ADD html.go ${APP}

WORKDIR ${APP}
RUN go get
RUN go build -o /goapp main.go html.go

ENTRYPOINT ["/goapp"]
EXPOSE 8000