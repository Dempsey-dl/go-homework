// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MetaNode is ERC20 {
    constructor() ERC20("MetaNodeToken","MetaNode") {
        _mint(msg.sender,1000*1_000_000_000_000_000_000);
    }

    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }
}
