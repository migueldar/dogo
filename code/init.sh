#!/bin/bash


case $1 in
	"start")
		airmon-ng start $2
		iwconfig "$2mon" channel 5.68G
		wireshark
		;;
	"stop")
		airmon-ng stop "$2mon"
		;;
	*)
		echo "usage ./init.sh <start | stop> <NIC name>"
		exit 1
		;;
esac
