FROM golang:1.23-rc-alpine3.19

WORKDIR /app

# コメントアウトを解除
COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

COPY ./backend .

CMD ["go", "run", "./cmd/api"]