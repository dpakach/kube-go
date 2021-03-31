FROM golang:1-alpine as build

RUN mkdir -p /app
WORKDIR /app
COPY main.go main.go
RUN go build main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main /app/main

EXPOSE 8080
CMD ["./main"]
