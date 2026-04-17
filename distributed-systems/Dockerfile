FROM registry.localhost:5000/public.ecr.aws/docker/library/golang:1.24-alpine AS builder

ARG APP_NAME

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ${APP_NAME} ./cmd/${APP_NAME}

FROM registry.localhost:5000/public.ecr.aws/ubuntu/ubuntu:24.04

ARG APP_NAME
ENV APP_NAME=${APP_NAME}

WORKDIR /app/${APP_NAME}

COPY --from=builder /app/${APP_NAME} .

EXPOSE 8080

CMD ./${APP_NAME}
