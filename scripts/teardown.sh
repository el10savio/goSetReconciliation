#!/usr/bin/env bash

docker ps -a | awk '$2 ~ /set/ {print $1}' | xargs -I {} docker rm -f {}
docker network rm set_network
