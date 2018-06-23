FROM golang:1.9-stretch AS build
WORKDIR /go/src/github.com/pollosp/pingdomcli/
RUN go get -d -v github.com/russellcardullo/go-pingdom/pingdom && go get -d -v github.com/urfave/cli
COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch
WORKDIR /root/
COPY --from=build /go/src/github.com/pollosp/pingdomcli/pingdomcli .
CMD ["./pingdomcli","-h"]
