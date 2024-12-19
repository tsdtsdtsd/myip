FROM busybox AS wgeter
FROM golang:latest AS builder

WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o myip .

FROM gcr.io/distroless/static-debian12
LABEL org.opencontainers.image.source="https://github.com/tsdtsdtsd/myip"
WORKDIR /opt/myip

COPY --from=builder /build/myip .
COPY --from=wgeter /bin/wget /bin/wget

ENTRYPOINT ["./myip"]