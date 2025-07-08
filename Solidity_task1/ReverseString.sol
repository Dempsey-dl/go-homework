// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
✅ 反转字符串 (Reverse String)
题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
*/

contract ReverseString {

    function ReverseStr(string memory symbol) public pure returns (string memory) {
        bytes memory strbytes = bytes(symbol);
        bytes memory restrbytes = new bytes(strbytes.length);

        for (uint i=0; i < restrbytes.length; i++) {
            restrbytes[i] = strbytes[restrbytes.length - i - 1];
        }

        return string(restrbytes);
    }
}