// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract Version {
    string private _version;

    constructor(string memory version_) {
        _version = version_;
    }

    function setVersion(
        string memory version_
    ) public returns (string memory) {
        _version = version_; 
        return _version;
    }

    function getVersion() public view returns (string memory) {
        return _version;
    }
}
