FROM golang:1

WORKDIR /go/src/app
COPY . .

CMD ["go", "run", "src/main.go"]
