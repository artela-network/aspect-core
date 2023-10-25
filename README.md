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

Artela SDK implement all of joinpoints.

**Warning**: This project is still in a very early stage, DO NOT USE IT IN PRODUCTION.


All of our major modules are currently under `/artela/aspect` folder, includes the following:


## Development
### Requires
*  [Go 1.20+](https://go.dev/dl)
*  [Pre-Commit 3.5+](https://pre-commit.com/)

The quality of project code needs to be checked through lint，please install the git hook scripts。
```shell
pre-commit install
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