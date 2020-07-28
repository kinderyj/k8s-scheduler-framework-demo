FROM debian:stretch-slim

WORKDIR /

COPY _output/bin/scheduler-framework-demo /usr/local/bin

CMD ["scheduler-framework-demo"]
