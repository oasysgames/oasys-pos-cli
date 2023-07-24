# oasys-pos-cli

Command Line Tool for manage proof-of-stake of Oasys Blockchain.

## Build

```shell
$ go get && go build -o oaspos
```

## Usage

```shell
$ oaspos --help 

Name:
  oaspos - Command Line Tool for manage proof-of-stake of Oasys Blockchain.

  Copyright 2022 Oasys | Blockchain for The Games All Rights Reserved.
  
Version:
  1.0.0

Usage:
  oaspos [command]

Available Commands:
  completion                  Generate the autocompletion script for the specified shell
  crypto:create-account       Create a new account.
  help                        Help about any command
  staker:claim-rewards        Withdraw staking rewards.
  staker:claim-unstakes       Withdraw unstaked tokens whose lock period has expired.
  staker:show-balance         Show token balance.
  staker:show-rewards         Show balance of staking rewards.
  staker:show-stakes          Show stake information.
  staker:stake                Stake tokens to validator.
  staker:unstake              Unstake tokens from validator.
  validator:activate          Change the validator status to active.
  validator:claim-commissions Withdraw validator commissions.
  validator:deactivate        Change the validator status to disable.
  validator:info              Show validator information.
  validator:info-all          Show all validator information.
  validator:info-slash        List of the number of slashings occurred in each recent epoch
  validator:join              Join as a validator in the proof-of-stake.
  validator:update-operator   Update the block signing address.

Flags:
  -h, --help   help for oaspos

Use "oaspos [command] --help" for more information about a command.
```

## Validator Subcommands

First, set the environment variable to the private key for sign the transaction.

```shell
$ export PRIVATE_KEY=0x0123456789abcdfe...
```

Usage:

```shell
# For mainnet
$ oaspos validator:[subcommand] --network mainnet [flags]

# For testnet
$ oaspos validator:[subcommand] --network testnet [flags]

# For custom
$ oaspos validator:[subcommand] --rpc https://example.com/ --chain-id 12345 [flags]
```
