"use strict"

const demoContractABI = require("./runner/generated/abi-contract.json");
const fs = require('fs')
const Web3 = require('web3');

const demoContractOptions = {
    data: demoContractABI.schedule.bytecode,
    gasPrice: 1000000010, // Default gasPrice set by Geth
    gas: 4000000
};
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

    // instantiate an instance of demo contract
    let schedule_contract = new web3.atl.Contract(demoContractABI.aspect.abi,
        web3.utils.aspectCoreAddr, demoAspectContractOptions);
    // deploy demo contract
    let schedule_instance = schedule_contract.deploy().send({from: accounts[0], nonce: nonceVal});
    let contractAddress="";
    schedule_contract = await schedule_instance.on('receipt', function (receipt) {
        console.log("=============== deployed contract ===============");
        console.log("contract address: " + receipt.contractAddress);
        console.log(receipt);
        contractAddress= receipt.contractAddress
    }).on('transactionHash', (txHash) => {
        console.log("deploy contract tx hash: ", txHash);
    });

    console.log("== ScheduleTo ==",contractAddress)

    console.log("== account ==",accounts[0])


    let nonceskip=6;
    for (;;){
        await new Promise(r => setTimeout(r, 5000));

        let result= await schedule_contract.methods.retrieve().call({from: accounts[0], nonce: nonceVal + nonceskip})

        console.log("==== schedule_contract result==="+ result);

        ++nonceskip
    }
}

f().then();
