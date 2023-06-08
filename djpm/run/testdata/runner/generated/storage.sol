// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Storage {
    uint256 number1;

    int32 number2;
    uint64 number3;
    string str1;

    bool bool1;

    // struct Person {
    //     uint64 id;
    //     uint32 balance;
    // }

    // mapping(string => Person) public accounts;

    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        number1 = number1 + num;
        number2 = 2000;
        // number3 = 3000;
        str1 = "hello";
        bool1 = true;
        // accounts["vv"] = Person(101, 9900);
    }

    /**
     * @dev Return value
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256) {
        return number1;
    }
}
