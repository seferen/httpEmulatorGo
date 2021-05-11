FROM golang:alpine
WORKDIR /app
COPY ./main.go /app
RUN go build -o /app/main main.go
ENTRYPOINT ["/app/main"]
EXPOSE 8080