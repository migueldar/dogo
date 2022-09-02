#!/bin/bash

case $1 in
    "up")
	docker-compose up --build -d
	docker exec -it attacker bash
	;;
    "down")
	docker-compose down
	;;
    "victim")
	docker exec -it victim bash
	;;
	"attacker")
	docker exec -it attacker bash
	;;
    *)
	echo "usage: ./start.sh [ up | down | victim | attacker ]"
	exit 1
	;;
esac
