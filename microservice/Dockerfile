FROM registry.localhost:5000/public.ecr.aws/docker/library/golang:1.24-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o microservice .

FROM registry.localhost:5000/public.ecr.aws/ubuntu/ubuntu:24.04

WORKDIR /app/microservice

COPY --from=builder /app/microservice .
EXPOSE 8080

CMD ["./microservice", "-port", "8080"]
