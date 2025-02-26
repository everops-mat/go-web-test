FROM golang:latest as builder
LABEL maintainer="Mat Kovach <mat.kovach@everops.com>"
WORKDIR /app
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest
RUN apk --no-cache add ca-certificates perl
WORKDIR /root/
COPY --from=builder /app/main .
COPY eo.pl .
RUN mkdir -p /tmp/both
EXPOSE 9991
CMD ["./main", "-wd", "/tmp/both", "-cmd", "/root/eo.pl"]
