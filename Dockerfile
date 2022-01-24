FROM golang:1.17.6

WORKDIR /arbitrage

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN go install -v /arbitrage/cmd/arbitrage

CMD ["arbitrage"]
