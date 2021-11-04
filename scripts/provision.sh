#!/bin/bash

# Initialize ENV Variables peers
# A list of all the ports
# of the set peers
declare -a peers=()
declare -a peer_id_list=()

# Name of the Set cluster network
# connecting the peers
network="set_network"

# number of peers to be provisioned
# Default 2 peers are provisioned
peers_count=$1

# Err check number of peers
# If no peers count is given default to 2
if [[ $peers_count -eq "" ]]; then
  peers_count=2
fi

echo "Number of peers: $peers_count"

# Exit when there are more than 1000 peers
if [[ $peers_count -ge 1000 ]]; then
  echo "Number of peers cannot be more than 1000"
  exit 255
fi

# Check if port is available and then
# append to peers starting from 8000
available_port=8000
provisioned_ports_count=0

echo "Cleaning previous stale peers"
docker ps -a | awk '$2 ~ /set/ {print $1}' | xargs -I {} docker rm -f {}
docker network rm "$network"

echo "Reserving ports for peers"

for port in {8000..9000}; do
  if [[ provisioned_ports_count -eq peers_count ]]; then
    break
  fi

  netstat -an | grep $port
  if [[ $? -ne 0 ]]; then
    peers+=($port)
    ((provisioned_ports_count++))
  fi
done

if [[ provisioned_ports_count -ne peers_count ]]; then
  echo "Unable to reserve ports for peers"
  exit 255
fi

echo "Reserved ports:" ${peers[*]}
comma_separated_peers=$(
  IFS=,
  echo "${peers[*]}"
)

# Docker create peers from peer list
# and pass PORT = peers[[i]]
echo "Provisioning Set Docker Cluster"

echo "Building Set Docker Image"
docker build -t set -f Dockerfile .

if [[ $? -ne 0 ]]; then
  echo "Unable To Build Set Docker Image"
  exit 255
fi

echo "Building Set Cluster Network"
docker network create "$network"

for ((id = 0; id < $peers_count; ++id)); do
  peer_id_list+=(peer-$id)
done

comma_separated_peer_id_list=$(
  IFS=,
  echo "${peer_id_list[*]}"
)

for peer_index in "${!peers[@]}"; do
  docker run -p "${peers[$peer_index]}":8080 --net $network -e "PEERS="$comma_separated_peer_id_list"" -e "NETWORK="$network"" -e="HOST=peer-$peer_index" --name="peer-$peer_index" -d set
done

# Docker list peers on success
echo "Set Cluster Nodes"
docker ps | grep 'set'
docker network ls | grep "$network"
