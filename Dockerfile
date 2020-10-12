FROM golang:1.15 as builder

RUN go env -w GO111MODULE=on

WORKDIR /project

COPY go.sum go.sum
COPY go.mod go.mod
COPY . .

RUN go mod download
RUN cd cmd && go build -o trots

FROM ubuntu:18.04

RUN mkdir /home/trots
RUN mkdir /tmp/reports

COPY --from=builder /project/cmd/trots /home/trots
COPY --from=builder /project/.docker/trots.yml /home/trots

WORKDIR /home/trots

ENTRYPOINT ["./trots"]
