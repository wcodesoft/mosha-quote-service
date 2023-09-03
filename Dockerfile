FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app/mosha-quote-service

COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

ENV HTTP_PORT 8280
ENV GRPC_PORT 8281
ENV MONGO_DB_HOST "mongodb://localhost:27017"
ENV AUTHOR_SERVICE_ADDRESS "localhost:8181"
ENV SENTRY_DSN ""
ENV SENTRY_SAMPLE_RATE "1.0"
ENV RELEASE_VERSION "dev"

WORKDIR /bin
COPY --from=builder /app/mosha-quote-service/app .

CMD ["./app"]
EXPOSE 8280 8281