FROM golang:latest AS builder

WORKDIR /build

ARG BUSYBOX_VERSION=1.35.0-i686-linux-musl
ADD https://busybox.net/downloads/binaries/$BUSYBOX_VERSION/busybox_WGET /wget
RUN chmod a+x /wget

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o myip .

FROM gcr.io/distroless/static-debian12
LABEL org.opencontainers.image.source="https://github.com/tsdtsdtsd/myip"
WORKDIR /opt/myip

COPY --from=builder /build/myip .
COPY --from=builder /wget /usr/bin/wget

ENTRYPOINT ["./myip"]