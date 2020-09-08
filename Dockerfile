FROM golang:latest
WORKDIR $GOPATH/src/satool
COPY . $GOPATH/src/satool
RUN go build app/satool/hangjia/hangjia.go
EXPOSE 8000 
ENTRYPOINT ["./satool"]