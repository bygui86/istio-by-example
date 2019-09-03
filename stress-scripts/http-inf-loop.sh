#!/bin/sh
# HTTPIE infinite loop

HOST="$1"
PORT=$2
ENDPOINT="$3"

# echo "$1 - $2 - $3"
# echo "HOST: $HOST"
# echo "PORT: $PORT"
# echo "ENDPOINT: $ENDPOINT"

while [ true ]
do
	http --pretty none --print b $HOST:$PORT/$ENDPOINT
	sleep 1
done
