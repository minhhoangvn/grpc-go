FROM golang:1.15.2-buster as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

RUN ls -la

FROM golang:1.15.2-buster as application

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 50001

RUN ls -la

CMD ["./main"]