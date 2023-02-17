FROM alpine:3.17.2
ARG TARGETOS
ARG TARGETARCH
COPY build/go-websocket-example_${TARGETOS}_${TARGETARCH} /usr/local/bin/go-websocket-example
RUN /usr/local/bin/go-websocket-example