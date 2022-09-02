FROM debian:bullseye-slim

COPY ./config/docker-entrypoint.sh /tmp/

RUN apt-get update  \
    && apt-get install -y net-tools tcpdump iputils-ping \
	&& chmod +x /tmp/docker-entrypoint.sh

CMD /tmp/docker-entrypoint.sh