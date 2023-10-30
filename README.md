<!--
parent:
  order: false
-->

<div align="center">
  <h1> Aspect Core </h1>
</div>

<div align="center">
  <a href="https://github.com/cosmos/cosmos-sdk/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/cosmos/cosmos-sdk.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/cosmos/cosmos-sdk">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/cosmos/cosmos-sdk" />
  </a>
  <a href="https://codecov.io/gh/cosmos/cosmos-sdk">
    <img alt="Code Coverage" src="https://codecov.io/gh/cosmos/cosmos-sdk/branch/main/graph/badge.svg" />
  </a>
</div>

We present Aspect Programming, a programming model for Artela Blockchain[1] that enables
native extensions on the blockchain. Aspect is the programmable extension used to dynamically
integrate custom functionality into the blockchain at runtime, working in conjunction with smart
contracts to enhance on-chain functionality. The distinguishing feature of Aspect is the ability
to access the system-level APIs of the base layer1
and perform designated actions at join points
throughout the transaction’s life cycle. Smart contracts can bind specified Aspects to activate
additional functionality. When a transaction invokes these smart contracts, it interacts with the
associated Aspects. Aspect Programming is also designed as a universal programming model applicable to any other modular blockchain execution layer2
. With Aspect Programming, developers
can implement basic logic in smart contracts and extend additional features in Aspects to build
feature-rich dApps beyond the capabilities of EVM[2].


*> *Warning**: This project is still in a very early stage, DO NOT USE IT IN PRODUCTION.

## Overview
This repository is the core code of the entire Aspect implementation. The main functions implemented include* Joinpoint
* JoinPoint
* Wasm Runner
* Wasm Host Api
* Chain Extension
  * Account abstraction
  * Schedule


## Development
### Requires
*  [Go 1.20+](https://go.dev/dl)
*  [Pre-Commit 3.5+](https://pre-commit.com/)

The quality of project code needs to be checked through lint，please install the git hook scripts。
```shell
pre-commit install
```

### Add Aspect
```go
// 1. init aspect Instance
// init Aspect
SetAspect(app.EvmKeeper.GetAspectCosmosProvider())


// 2.Add `joinpoint` code to the appropriate location, such as begin block
import "github.com/artela-network/aspect-core/djpm"

request := &asptypes.EthBlockAspect{Header: header, GasInfo: &asptypes.GasInfo{...}}
// do aspect
receive := djpm.AspectInstance().OnBlockFinalize(request)
hasErr, receiveErr := receive.HasErr()
if hasErr {
ctx.Logger().Error("Aspect.OnBlockFinalize Return Error ", receiveErr.Error(), "height", request.Header.Number)
}
...

```

### Directory Structure

```shell
├── chaincoreext               // Chain scalability 
│   ├── account_abstraction    // Implementing ethereum `Account Abstraction`
│   ├── jit_inherent           // Initiate message calls through account abstraction
│   └── scheduler              // Provides the functionalities to generate automated on-chain transaction
├── djpm                       // Manager for all aspects
│   ├── contract               // Call smart contract related functions
│   └── run                    // Provides basic aspect functionalities, including call WASM runtime, gas metering and etc.
├── integration                // Message calls interface provider
├── proto                      // Define protobuf message for Aspect
│   └── message
├── scripts                     
└── types                      // Basic data type of aspects

```


## License
Copyright © Artela Network, Inc. All rights reserved.

Licensed under the [Apache v2](LICENSE) License.