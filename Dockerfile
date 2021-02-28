FROM golang:1.14

RUN mkdir -p /go/src/github.com/yuki9541134/randish-api

WORKDIR /go/src/github.com/yuki9541134/randish-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "cmd/main.go"]
