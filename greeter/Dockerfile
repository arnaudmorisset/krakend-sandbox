FROM golang:1.22.0-bullseye AS build
WORKDIR /app
COPY . .
RUN go build

FROM debian:stable-slim
WORKDIR /app
COPY --from=build /app/greeter ./
EXPOSE 50051 50051
CMD ["./greeter"]
