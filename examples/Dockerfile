FROM golang

RUN mkdir bin

RUN mkdir src/microservice
COPY microservice.go $GOPATH/src/microservice/
RUN go install microservice
RUN chmod +x $GOPATH/bin/microservice

EXPOSE 3100

ENTRYPOINT ["/go/bin/microservice"]