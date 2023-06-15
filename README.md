<!--
parent:
  order: false
-->

<div align="center">
  <h1> Artela SDK </h1>
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

**Note**: Requires [Go 1.19+](https://go.dev/dl)

## System Overview

All of our major modules are currently under `/artela/aspect` folder, includes the following:
- 
- `djpm` : Manager for all aspects
- `djpm/run` : Provides basic aspect functionalities, including call WASM runtime, gas metering and etc.
- `scheduler` : Provides the functionalities to generate automated on-chain transaction
- `proto` : define protobuf message for Aspect
- `types` : Basic data types of aspects