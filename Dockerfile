FROM golang:1.21.5-bullseye

WORKDIR /app
COPY . .

RUN go get
RUN go build -o bin .

ENTRYPOINT ["/app/bin"]