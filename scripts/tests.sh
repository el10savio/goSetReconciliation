#!/usr/bin/env bash

echo "Provisioning Cluster With 2 Nodes"
scripts_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

bash ${scripts_dir}/provision.sh 2

echo "Cluster Sanity Tests"
bats ${scripts_dir}/bats/cluster-sanity.bats
