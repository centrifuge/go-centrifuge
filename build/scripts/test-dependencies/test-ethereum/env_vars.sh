#!/usr/bin/env bash

RPC_PORT=9545
WS_PORT=9546
NETWORK_ID=1337
IDENTITY=CentTestEth
GETH_DOCKER_CONTAINER_NAME="geth-node"
CENT_ETHEREUM_CONTRACTS_DIR=${PARENT_DIR}/vendor/github.com/centrifuge/centrifuge-ethereum-contracts
CENT_ETHEREUM_CONTEXTWAITTIMEOUT="180s"
CENT_ETHEREUM_NODEURL=${CENT_ETHEREUM_NODEURL:-http://localhost:$RPC_PORT}
CENT_ETHEREUM_GASLIMIT=4712388
CENT_ETHEREUM_GASPRICE=1000000000
CENT_ETHEREUM_GETH_START_TIMEOUT=${CENT_ETHEREUM_GETH_START_TIMEOUT_OVERRIDE:-600} # In Seconds, default 10 minutes
CENT_ETHEREUM_GETH_START_INTERVAL=${CENT_ETHEREUM_GETH_START_INTERVAL_OVERRIDE:-2} # In Seconds, default 2 seconds
