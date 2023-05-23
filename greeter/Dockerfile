FROM golang:1.20.3-bullseye AS build
WORKDIR /app
COPY . .
RUN go build

FROM debian:stable-slim
COPY --from=build . .
WORKDIR /app
EXPOSE 50051 50051
CMD ["./greeter"]
