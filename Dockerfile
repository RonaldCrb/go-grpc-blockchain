FROM golang:1.13.8-alpine3.11 as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /build/blochain-server server/main.go 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /build/blochain-client client/main.go

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/blochain-server .
COPY --from=builder /build/blochain-client .

# arguments that can be overridden

CMD ["./blockchain-server"]
