FROM golang:latest as builder

WORKDIR /workdir

COPY  . .

RUN go build .

FROM golang:latest

WORKDIR /workdir

COPY --from=builder /workdir/http-go-mux .

COPY entrypoint.sh .

ENTRYPOINT /workdir/entrypoint.sh
