// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.2 <0.9.0;

contract AspectCore {

    function deploy(
        bytes calldata code,
        KVPair[] calldata properties,
        address account,
        bytes calldata proof,
        uint256 joinPoints
    ) public {}

    function upgrade(
        address aspectId,
        bytes calldata code,
        KVPair[] calldata properties,
        uint256 joinPoints
    ) public {}

    function bind(
        address aspectId,
        uint256 aspectVersion,
        address accountAddr,
        int8 priority
    ) public {}

    function unbind(address aspectId, address accountAddr) public {}

    function changeVersion(
        address aspectId,
        address accountAddr,
        uint64 version
    ) public {}

    function versionOf(address aspectId) public view returns (uint64 version) {}

    function aspectsOf(address accountAddr)
    public
    view
    returns (AspectBoundInfoArr[] memory aspectBoundInfo)
    {}

    function boundAddressesOf(address aspectId)
    public
    view
    returns (address[] memory accounts)
    {}

    function entrypoint(address aspectId, bytes calldata data)
    public
    view
    returns (bytes memory result)
    {}

    struct KVPair {
        string key;
        bytes value;
    }
    struct AspectBoundInfoArr {
        address aspectId;
        uint64 version;
        int8 priority;
    }
    /*
    function getAspect(address aspectId) public view returns (Aspect memory aspect) {}

    struct Aspect {
        address aspectId;
        uint64  aspectVersion;
        uint64[]  allVersions;
        bytes  aspectCode;
        KVPair[]  properties;
        string[]  joinPoints;
        bytes  proof;
        address account;
    }
    */
}
