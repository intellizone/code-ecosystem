FROM public.ecr.aws/docker/library/golang:1.24-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o readinglist ./cmd/api/main.go

FROM public.ecr.aws/ubuntu/ubuntu:24.04

WORKDIR /app/readinglist

COPY --from=builder /app/readinglist .

EXPOSE 8080

CMD ["./readinglist", "-port", "8080"]