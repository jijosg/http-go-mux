FROM golang:latest as builder

WORKDIR /workdir

COPY  . .

RUN go build .

FROM golang:latest

WORKDIR /workdir

COPY --from=builder /workdir/http-sample .
COPY --from=builder /workdir/data.db .

COPY entrypoint.sh .


ENTRYPOINT /workdir/entrypoint.sh
