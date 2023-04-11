#!/bin/sh

set -e

host="$1"
shift
port="$1"
shift
cmd="$@"

until nc -z -v -w30 "$host" "$port"; do
  echo "Waiting for $host:$port..."
  sleep 1
done

echo "$host:$port is available. Starting the application..."
exec $cmd
