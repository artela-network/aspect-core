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
    let schedule_contract = new web3.atl.Contract(demoContractABI.schedule.abi,
        web3.utils.aspectCoreAddr, demoContractOptions);
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



    // instantiate an instance of demo contract
    let contract = new web3.atl.Contract(demoContractABI.aspect.abi,
        web3.utils.aspectCoreAddr, demoAspectContractOptions);
    // deploy demo contract
    let instance = contract.deploy().send({from: accounts[0], nonce: nonceVal+1});
    contract = await instance.on('receipt', function (receipt) {
        console.log("=============== deployed contract ===============");
        console.log("contract address: " + receipt.contractAddress);
        console.log(receipt);
    }).on('transactionHash', (txHash) => {
        console.log("deploy contract tx hash: ", txHash);
    });


    // load aspect code and deploy
    let aspectCode = fs.readFileSync('./build/release.wasm', {
        encoding: "hex"
    });
    // instantiate an instance of aspect
    let aspect = new web3.atl.Aspect(
        web3.utils.aspectCoreAddr, demoContractOptions);
    instance = aspect.deploy({
        data: '0x' + aspectCode,
        properties: [{'key': '0x00', 'value': '0x02'},{'key': '0x032322', 'value': '0x2221'},{'key': '0x8a5363686564756c65546f', 'value': contractAddress}]
    }).send({from: accounts[0], nonce: nonceVal + 2});

    aspect = await instance.on('receipt', (receipt) => {
        console.log("=============== deployed aspect ===============");
        console.log("aspect address: " + aspect.options.address);
        console.log(receipt);
    }).on('transactionHash', (txHash) => {
        console.log("deploy aspect tx hash: ", txHash);
    });

    // bind the smart contract with aspect
    await contract.bind({
        priority: 1,
        aspectId: aspect.options.address,
        aspectVersion: 1,
    }).send({from: accounts[0], nonce: nonceVal + 3})
        .on('receipt', function (receipt) {
            console.log("=============== bind aspect ===============")
            console.log(receipt)
        })
        .on('transactionHash', (txHash) => {
            console.log("contract binding tx hash: ", txHash);
        });

    await new Promise(r => setTimeout(r, 5000));

    // call the smart contract, aspect should be triggered
    await contract.methods.store(100)
        .send({from: accounts[0], nonce: nonceVal + 4})
        .on('receipt', (receipt) => {
            console.log("=============== called store ===============")
            console.log(receipt);
        })
        .on('transactionHash', (txHash) => {
            console.log("call contract tx hash: ", txHash);
        });

    let result= await contract.methods.retrieve().call({from: accounts[0], nonce: nonceVal + 5})
    console.log("==== reuslt==="+ result);

}

f().then();
