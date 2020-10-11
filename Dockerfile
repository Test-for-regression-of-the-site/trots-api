FROM golang:1.15 as builder

RUN go env -w GO111MODULE=on

WORKDIR /project

COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download

COPY . .
RUN cd cmd && go build -o trots

FROM ubuntu:18.04

RUN mkdir /opt/trots
COPY --from=builder /project/cmd/trots /opt/trots
COPY --from=builder /project/trots.yml /opt/trots
RUN chmod +x /opt/trots/trots
WORKDIR /opt/trots

ENTRYPOINT ["./trots"]
