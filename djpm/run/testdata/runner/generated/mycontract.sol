// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
 
contract MyContract {
    mapping(address => uint) public balances;

		mapping(address => string) public memo;

    function deposit(string calldata user) public payable {
        balances[msg.sender] += msg.value;
				memo[msg.sender] = user;
    }
 
    function withdraw() public {
        uint bal = balances[msg.sender];
        require(bal > 0);
 
        (bool sent,) = msg.sender.call{value: bal}(""); // Vulnerability of re-entrancy
        require(sent, "Failed to send Ether");
 
        balances[msg.sender] = 0;
    }
 
    // Helper function to check the balance of this contract
    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}