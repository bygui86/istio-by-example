#!/bin/bash
# Kubernetes health check

# set -e

HOST="$1"
PORT=$2
ENDPOINT="$3"

RESPONSE=$(curl --connect-timeout 2 --max-time 2 --write-out %{http_code} --silent --output /dev/null $HOST:$PORT/$ENDPOINT)
if [[ $RESPONSE != "200" ]]; then
	echo "KO :("
	exit -1
else
	echo "OK :)"
fi
