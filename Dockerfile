FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go build -o datetimeapp

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/datetimeapp .
EXPOSE 8080
CMD ["./datetimeapp"]
