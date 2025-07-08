// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
✅  用 solidity 实现罗马数字转数整数
题目描述在 https://leetcode.cn/problems/integer-to-roman/description/
*/
contract RomanToint {

        // 罗马字符到数字的映射
        mapping(bytes1 => uint256) romanValues;
        
        constructor() {
        romanValues['I'] = 1;
        romanValues['V'] = 5;
        romanValues['X'] = 10;
        romanValues['L'] = 50;
        romanValues['C'] = 100;
        romanValues['D'] = 500;
        romanValues['M'] = 1000;
        }


    function RomanInt(string memory s) public view returns (uint256) {
        bytes memory roman = bytes(s);
        uint256 total = 0;
        uint256 prevValue = 0;
        for (uint i = roman.length; i > 0; i--) {
            uint256 curvalue = romanValues[roman[i - 1]];
            if (curvalue < prevValue) {
                total -= curvalue;
            } else {
                total += curvalue;
            }
            prevValue = curvalue;
        }
        return total;
    }   
}