#!/bin/sh
# CURL infinite loop

HOST="$1"
PORT=$2
ENDPOINT="$3"

# echo "$1 - $2 - $3"
# echo "HOST: $HOST"
# echo "PORT: $PORT"
# echo "ENDPOINT: $ENDPOINT"

while [ true ]
do
	curl -s -L -I -w "%{http_code}" -o /dev/null --connect-timeout 3 --max-time 3 $HOST:$PORT/$ENDPOINT
	printf "\n"
	sleep 1
done
