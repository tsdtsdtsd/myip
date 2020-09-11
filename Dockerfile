FROM golang:latest as builder

WORKDIR /go/src/app/
COPY . .

RUN go mod download
#RUN go build -v -o myip
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o myip .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/myip .
CMD ["./myip"]
