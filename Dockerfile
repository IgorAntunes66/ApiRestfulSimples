FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /tasklist ./cmd/api

FROM scratch

COPY --from=builder /tasklist /tasklist

EXPOSE 8080

CMD [ "/tasklist" ]