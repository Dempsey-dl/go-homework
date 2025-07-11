// SPDX-License-Identifier: MIT
pragma solidity ^0.8;



/*
✅ 创建一个名为Voting的合约，包含以下功能：
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    // //一个mapping来存储候选人的得票数
    // mapping (address => uint) private condidateVotes;
    // address[] condidates;
    // //一个vote函数，允许用户投票给某个候选人
    // function vote(address user) public {
    //     if(condidateVotes[user] == 0)
    //     {
    //         condidates.push(user);
    //     }
    //     condidateVotes[user] += 1;
    // }
    // //一个getVotes函数，返回某个候选人的得票数
    // function getVotes(address name) public view returns (uint) {
    //     return condidateVotes[name];
    // }
    // //一个resetVotes函数，重置所有候选人的得票数
    // function resetVotes() public {
    //     uint  len = condidates.length;
    //     for (uint i=0; i < len; i++) {
    //         delete condidateVotes[condidates[i]];
    //     }
    // }


    //一个mapping来存储候选人的得票数
    mapping (string => uint) private condidateVotes;
    string[] condidates;
    //一个vote函数，允许用户投票给某个候选人
    function vote(string memory user) public {
        if(condidateVotes[user] == 0)
        {
            condidates.push(user);
        }
        condidateVotes[user] += 1;
    }
    //一个getVotes函数，返回某个候选人的得票数
    function getVotes(string memory name) public view returns (uint) {
        return condidateVotes[name];
    }
    //一个resetVotes函数，重置所有候选人的得票数
    function resetVotes() public {
        uint  len = condidates.length;
        for (uint i=0; i < len; i++) {
            delete condidateVotes[condidates[i]];
        }
    }
}