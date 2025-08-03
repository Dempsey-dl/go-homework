// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.20;

import "./MetaNodeStake.sol";

contract MetaNodeStakeV2 is MetaNodeStake1 {
    // 预留存储槽防止冲突
    uint256[50] private __gap;

    function currentVerSion() public pure override returns (string memory) {
        return "V3.0";
    }
    
    // 新增获取合约余额函数
    function getContractBalance() public view returns (uint256) {
        return address(this).balance;
    }
}
