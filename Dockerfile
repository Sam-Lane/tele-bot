FROM golang

WORKDIR $GOPATH/src/github.com/sam-lane/tele-bot


COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["./tele-bot"]
