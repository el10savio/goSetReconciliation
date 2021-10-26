#!/usr/bin/env bash

echo "Provisioning Cluster With 2 Nodes"
scripts_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
bash ${scripts_dir}/provision.sh 2 > /dev/null

echo "Cluster Sanity Tests"
bats --tap ${scripts_dir}/bats/cluster-sanity.bats

echo "Full Sync Tests"
bats --tap ${scripts_dir}/bats/full-sync.bats
bats --tap ${scripts_dir}/bats/full-sync-other-node.bats

echo "Mixed Sync Tests"
bats --tap ${scripts_dir}/bats/mixed-sync.bats

echo "Resync Tests"
bats --tap ${scripts_dir}/bats/resync.bats

echo "Tearing Down Cluster"
bash ${scripts_dir}/teardown.sh > /dev/null
