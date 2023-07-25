FROM golang:1.20-alpine as builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/api

FROM alpine:3.18
COPY --from=builder main /bin/main
EXPOSE 8081
ENTRYPOINT ["/bin/main"]