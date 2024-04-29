FROM golang:1.22 as builder

ENV GOPROXY=direct

WORKDIR /app


COPY go.mod ./
COPY go.sum ./


RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .


EXPOSE 8080

CMD ["./main"]