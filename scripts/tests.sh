#!/usr/bin/env bash

echo "Provisioning Cluster With 2 Nodes"
scripts_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Cluster Sanity Tests"
bash ${scripts_dir}/provision.sh 2
bats --tap ${scripts_dir}/bats/cluster-sanity.bats
bash ${scripts_dir}/teardown.sh

echo "Full Sync Tests"
bash ${scripts_dir}/provision.sh 2
bats --tap ${scripts_dir}/bats/full-sync.bats
bash ${scripts_dir}/teardown.sh

echo "Mixed Sync Tests"
bash ${scripts_dir}/provision.sh 2
bats --tap ${scripts_dir}/bats/mixed-sync.bats
bash ${scripts_dir}/teardown.sh
