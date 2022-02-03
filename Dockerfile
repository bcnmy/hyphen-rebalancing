FROM golang:1.17.6

WORKDIR /rebalancing

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN go install -v /rebalancing/cmd/rebalancing

CMD ["rebalancing"]
