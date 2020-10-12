FROM golang:1.14-alpine as builder

LABEL maintainer="Dušan Simić <dusan.simic1810@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

ENV PORT=3000
ENV GIN_MODE=release

EXPOSE ${PORT}

CMD ["./main"]
