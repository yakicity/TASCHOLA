FROM golang:1.19

WORKDIR /app/backend
COPY ./backend .

RUN go mod download
RUN go mod tidy

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air
RUN go get -u github.com/rubenv/sql-migrate/...
