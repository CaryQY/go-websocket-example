FROM scratch
ARG TARGETOS
ARG TARGETARCH
COPY build/go-websocket-example_${TARGETOS}_${TARGETARCH} /usr/local/bin/go-websocket-example
CMD ["/usr/local/bin/go-websocket-example"]