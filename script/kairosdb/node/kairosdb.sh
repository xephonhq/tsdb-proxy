#!/usr/bin/env bash

echo "waiting for cassandra to start"
# FIXME： timeout is not included
wait-for-it cassandra:9042
echo "cassandra started"
/opt/kairosdb/bin/kairosdb.sh run
