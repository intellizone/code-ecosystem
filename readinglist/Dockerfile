FROM registry.localhost:5000/public.ecr.aws/docker/library/golang:1.24-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o readinglist ./cmd/api
RUN go build -o readinglist-ui ./cmd/web

FROM registry.localhost:5000/public.ecr.aws/ubuntu/ubuntu:24.04

WORKDIR /app/readinglist

COPY --from=builder /app/readinglist .
COPY --from=builder /app/readinglist-ui .

EXPOSE 8080

CMD ["./readinglist", "-port", "8080"]
