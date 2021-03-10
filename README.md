# Centrifuge OS node 

[![Tests](https://github.com/centrifuge/go-centrifuge/actions/workflows/tests.yml/badge.svg?branch=develop)](https://github.com/centrifuge/go-centrifuge/actions/workflows/tests.yml)
[![GoDoc Reference](https://godoc.org/github.com/centrifuge/go-centrifuge?status.svg)](https://godoc.org/github.com/centrifuge/go-centrifuge)
[![codecov](https://codecov.io/gh/centrifuge/go-centrifuge/branch/develop/graph/badge.svg)](https://codecov.io/gh/centrifuge/go-centrifuge)
[![Go Report Card](https://goreportcard.com/badge/github.com/centrifuge/go-centrifuge)](https://goreportcard.com/report/github.com/centrifuge/go-centrifuge)

`go-centrifuge` is the go implementation of the Centrifuge OS interacting with the peer to peer network, Centrifuge Chain, and our Ethereum smart contracts. 

**Getting help:** Head over to our developer documentation at [developer.centrifuge.io](http://developer.centrifuge.io) to learn how to setup a node and interact with it. If you have any questions, feel free to join our [discord](https://centrifuge.io/discord) 

**DISCLAIMER:** The code released here presents a very early alpha version that should not be used in production and has not been audited. Use this at your own risk.

## Pre-requisites
- Go >= 1.15.x
- Nodejs 10.x.x
- Truffle 5.1.29
- Dapp tools

## Running tests
There 4 different flavours of tests in the project
- Unit tests(unit)
- Command-line tests(cmd)
- Integration tests(integration)
- Environment tests(testworld): spins up multiple go-centrifuge nodes and local ethereum and centrifuge chains

To run all the test flavours: `make run-tests`

To run specific test flavour: `test=unit make run-tests`

To force ethereum smart contracts to be deployed again: `FORCE_MIGRATE='true' test=cmd make run-tests`

Note: `unit` tests doesn't require any smart contract deployments and when run with only `unit` flavour, smart contracts are not deployed.

## Installation
To install, run `make install` will compile project to binary `centrifuge` and be placed under `GOBIN`.

Ensure `GOBIN` is under `PATH` to call the binary globally.

## Spin-up local test environment:
For development, we use Docker Compose locally

### Centrifuge Chain
Local Centrifuge Chain comes with a set of preconfigured accounts to be used.

Start local centrifuge chain via `./build/scripts/docker/run.sh ccdev`

For more info: https://github.com/centrifuge/centrifuge-chain

### Run a Geth node locally in dev mode

Start the local geth node via `./build/scripts/docker/run.sh dev`

By default, geth node uses:
- ETH_DATADIR=${HOME}/Library/Ethereum
- RPC_PORT=9545
- WS_PORT=9546

## API definitions
Node APIs are published to swagger hub. 
For the latest APIs, please see here: [APIs](https://app.swaggerhub.com/apis/centrifuge.io/cent-node/)

