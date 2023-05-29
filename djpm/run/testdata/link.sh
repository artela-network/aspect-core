#!/bin/bash

web3jsDir=/Users/admin/mytech/go-work/src/github.com/artela-network/web3.js/packages
currentDir=$(pwd)

packages=()
for d in "$web3jsDir/*"
do
  packages+=($(basename $d))
done
npm link "${packages[@]}"
