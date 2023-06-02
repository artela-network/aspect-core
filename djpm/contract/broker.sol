interface Broker {
    // scheduler will call this interface to check how many
    // tokens that the given Aspect can charge,
    // unit of return value is Wei
    function allowance(address aspectId)  external view returns (uint256 valueWei);
}