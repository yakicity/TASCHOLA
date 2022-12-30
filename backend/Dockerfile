FROM golang:1.19

WORKDIR /app/backend
COPY ./backend .

# wait-for-it.sh is a script that waits for a service to be available
# before starting the server. This is useful for when the database
# container is started before the server container.
COPY ./wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

RUN go mod download
RUN go mod tidy

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air
