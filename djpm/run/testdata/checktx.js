"use strict"

const demoContractABI = require("./runner/generated/abi-contract.json");
const fs = require('fs')
const Web3 = require('web3');


const demoAspectContractOptions = {
    data: demoContractABI.aspect.bytecode,
    gasPrice: 1000000010, // Default gasPrice set by Geth
    gas: 4000000
};

async function f() {
    // init connection to Artela node
    const web3 = new Web3('http://127.0.0.1:8545');

    // retrieve accounts
    let accounts = await web3.eth.getAccounts();

    // retrieve current nonce
    let nonceVal = await web3.eth.getTransactionCount(accounts[0]);

    //todo replace contract address
    let contract_addr="0x4E83763CBb8c553aA8E9F49d1E73Fc3C9377600e";

    // instantiate an instance of demo contract
    let schedule_contract = new web3.atl.Contract(demoContractABI.aspect.abi,
        contract_addr, demoAspectContractOptions);


    let nonceskip=1000;
    for (;;){
        await new Promise(r => setTimeout(r, 5000));

        let result= await schedule_contract.methods.retrieve().call({from: accounts[0], nonce: nonceVal + nonceskip})

        console.log("==== schedule_contract result==="+ result);

        ++nonceskip
    }
}

f().then();
