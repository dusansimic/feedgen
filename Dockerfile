FROM golang:1.14-alpine as builder

LABEL maintainer="Dušan Simić <dusan.simic1810@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server/server ./server

FROM golang:1.14-alpine

WORKDIR /app

COPY --from=builder /app/server/server .

ENV PORT=3000
ENV GIN_MODE=release

EXPOSE ${PORT}

CMD ["/app/server"]
