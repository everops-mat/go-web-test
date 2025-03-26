FROM golang:1.23 as builder
LABEL maintainer="Mat Kovach <mat.kovach@everops.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o eo-sayings ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/eo-sayings /root/eo-sayings
COPY config/sayings.txt /root/config/sayings.txt
EXPOSE 8080
CMD ["/root/eo-sayings"]
