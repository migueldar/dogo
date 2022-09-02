FROM debian:bullseye-slim

RUN apt-get update \
    && apt-get -y install golang libpcap-dev ca-certificates net-tools tcpdump iputils-ping

WORKDIR /code

CMD tail -f /dev/null
