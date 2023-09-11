package account_abstraction

// 1. brew tap ethereum/ethereum
// 2. brew install solidity
//go:generate solc contract/contracts/interfaces/IEntryPoint.sol --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ./ --overwrite
//go:generate abigen --pkg account_abstraction --out entrypoint.go --combined-json ./combined.json
