FROM golang:latest as builder

WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o myip .

FROM gcr.io/distroless/static-debian12
WORKDIR /opt/myip
COPY --from=builder /build/myip .
CMD ["./myip"]
