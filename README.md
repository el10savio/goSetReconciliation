# goSetReconciliation

An implementation to sync distributed sets using bloom filters. Based on the paper "Low complexity set reconciliation using Bloom filters" by Magnus Skjegstad and Torleiv Maseng.

## Introduction

Syncing multiple distributed sets over a distributed system and hard and tedious. To tackle this the paper aims at syncing multiple lists by sending across bloom filters to each node to then calculate the elements missing between the sets. This idea makes syncing sets much easier and of lower complexity.

## Example

```
$ curl -i -X POST localhost:8000/set/add -d {"values": [1,2,6]}
$ curl -i -X POST localhost:8001/set/add -d {"values": [1,2,3,4,5]}
$ curl -i -X GET localhost:8000/set/sync
$ curl -i -X GET localhost:8000/set/list > [1,2,6,3,4,5]
$ curl -i -X GET localhost:8000/set/list > [1,2,3,4,5,6]
```

## Steps

To provision the cluster:

```
$ git clone https://github.com/el10savio/goSetReconciliation
$ cd goSetReconciliation
$ make provision
```

This creates a 2 node Set cluster established in their own docker network.

To view the status of the cluster

```
$ make info
```

This provides information on the cluster and its associated ports to access each node. An example of the output seen in `make info` would be like below:

```
d3fd26dd4df3  set  "/go/bin/set"  2 hours ago  Up 2 hours  0.0.0.0:8004->8080/tcp  peer-1
8830feb6cd68  set  "/go/bin/set"  2 hours ago  Up 2 hours  0.0.0.0:8003->8080/tcp  peer-0
```


Now we can also send requests to add, list, and sync values to any peer node using its port allocated.

```
$ curl -i -X POST localhost:<peer-port>/set/add -d {"values": <values>}
$ curl -i -X GET localhost:<peer-port>/set/list
$ curl -i -X GET localhost:<peer-port>/set/sync
```

In the logs for each peer docker container, we can see the logs of the peer nodes getting in sync when issuing the `/sync` request.

To tear down the cluster and remove the built docker images:

```
$ make clean
```

This is not certain to clean up all the locally created docker images at times. You can do a docker rmi to delete them.

## Testing

To provision the cluster and run automated end to end tests you can use `make e2e`. This uses BATS bash testing to run curl requests to each node and asserts the output received.

```
$ make e2e
Running E2E Testing On Set Cluster
bash scripts/tests.sh
Provisioning Cluster With 2 Nodes
Error: No such network: set_network
Cluster Sanity Tests
1..6
ok 1 Check Replicas Count
ok 2 Check Replicas Are Avaialable
ok 3 Writes Are Succesfull
ok 4 Reads Are Succesfull
ok 5 Writes Are Idempotent
ok 6 Set Debug Clear
Full Sync Tests
1..2
ok 1 Add Elements To One Node Only & Check For Successfull Sync
ok 2 Set Debug Clear
1..2
ok 1 Add Elements To One Node Only & Check For Successfull Sync From Other Node
ok 2 Set Debug Clear
Mixed Sync Tests
1..2
ok 1 Add Different Elements To Nodes & Check For Successfull Sync
ok 2 Set Debug Clear
Tearing Down Cluster
```

## References

- [ Low complexity set reconciliation using Bloom filters ](https://dl.acm.org/doi/10.1145/1998476.1998483) [Magnus Skjegstad and Torleiv Maseng]
- [ Low Complexity Set Reconciliation Using Bloom Filters | Paper Review ](https://www.youtube.com/watch?v=xuddEiu-t-8) [Heidi Howard]
