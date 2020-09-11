FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -v

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/app/myip .
CMD ["./myip"]