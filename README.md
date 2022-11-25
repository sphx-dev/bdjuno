# BDJuno

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/forbole/bdjuno/Tests)](https://github.com/forbole/bdjuno/actions?query=workflow%3ATests)
[![Go Report Card](https://goreportcard.com/badge/github.com/forbole/bdjuno)](https://goreportcard.com/report/github.com/forbole/bdjuno)
![Codecov branch](https://img.shields.io/codecov/c/github/forbole/bdjuno/cosmos/v0.40.x)

BDJuno (shorthand for BigDipper Juno) is the [Juno](https://github.com/forbole/juno) implementation
for [BigDipper](https://github.com/forbole/big-dipper).

It extends the custom Juno behavior by adding different handlers and custom operations to make it easier for BigDipper
showing the data inside the UI.

All the chains' data that are queried from the RPC and gRPC endpoints are stored inside
a [PostgreSQL](https://www.postgresql.org/) database on top of which [GraphQL](https://graphql.org/) APIs can then be
created using [Hasura](https://hasura.io/).

## Usage

To know how to setup and run BDJuno, please refer to
the [docs website](https://docs.bigdipper.live/cosmos-based/parser/overview/).

## Testing

If you want to test the code, you can do so by running

```shell
$ make test-unit
```

**Note**: Requires [Docker](https://docker.com).

This will:

1. Create a Docker container running a PostgreSQL database.
2. Run all the tests using that database as support.

## Local launch in docker compose

1. Create `.bdjuno` directory in the root of the repo and put there next two files:
   1. `config.yaml` which you can copy from `config-sample.yaml`. Do next change there:
      1. Replace `YOUR_CHAIN_ACC_PREFIX` with your account prefix(`coredev` should work for you). This will let you have separate accounts with the same private keys for different chains.
      2. Replace `YOUR_NODE_IP` with `cored` address.
   2. `genesis.json` which you can find in two ways:
      1. At cored `/genesis` endpoint(`http://127.0.0.1:26557/genesis?` for example)
      2. At mounted directory `~/.cache/crust/znet/znet/app/coredev-00/coreum-devnet-1/config`
   
2. Build and start
```bash
docker-compose up
```

* Open [hasura UI](http://localhost:8080/console) and check that it works correctly.
The password is defined in docker-compose.yaml and set to "myadminsecretkey" by default.

### Remarks

In case you run the bdjuno with the connection to old-running node you might face the error in the logs
```
error while getting staking pool: rpc error: code = Internal desc = UnmarshalJSON cannot decode empty bytes
```

This is expected since the node doesn't store all staking pool for all heights. 
