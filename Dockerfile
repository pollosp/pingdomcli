FROM golang:1.9-stretch AS builder
WORKDIR /go/src/github.com/pollosp/pingdomcli/
RUN go get -d -v github.com/russellcardullo/go-pingdom/pingdom && go get -d -v github.com/urfave/cli
RUN useradd -u 10001 scratchuser
COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app/
COPY --from=builder /go/src/github.com/pollosp/pingdomcli/pingdomcli .
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
CMD ["./pingdomcli","-h"]
