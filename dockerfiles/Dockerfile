FROM ugwuanyi/golang-python:1.17-3.9 as builder

WORKDIR /build

COPY . .

RUN make build

FROM ugwuanyi/golang-python:1.17-3.9

COPY --from=builder /build/dec /usr/local/bin/

ENTRYPOINT [ "dec" ]
