// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Storage {

    uint256 number;


    function isOwner(address sender)  public view returns  (bool){

        if((sender=="0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2")||(sender=="0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db")){
        return true;
        }
        return false;

    }

    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        number = num;
    }

    /**
     * @dev Return value
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256){
        return number;
    }
}