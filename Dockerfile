FROM scratch 
ARG TARGETOS
ARG TARGETARCH
COPY build/go-websocket-example_${TARGETOS}_${TARGETARCH} /opt/app/gows
CMD ["/opt/app/gows"]
