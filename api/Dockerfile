FROM golang:1.22.4-bookworm AS builder
WORKDIR /opt/build
COPY . .
RUN curl -LO https://github.com/tinygo-org/tinygo/releases/download/v0.32.0/tinygo_0.32.0_amd64.deb && dpkg -i tinygo_0.32.0_amd64.deb
RUN tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm main.go

FROM scratch
COPY --from=builder /opt/build/spin.toml .
COPY --from=builder /opt/build/main.wasm .
