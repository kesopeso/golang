FROM golang:alpine AS build
WORKDIR /app
COPY ./go.mod ./server.go ./
RUN go build -o ./server

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]
