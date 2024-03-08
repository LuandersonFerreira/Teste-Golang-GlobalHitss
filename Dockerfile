FROM golang:1.22 as builder
WORKDIR /app
COPY . .


RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/api ./
EXPOSE 9000
ENTRYPOINT ["./api"]