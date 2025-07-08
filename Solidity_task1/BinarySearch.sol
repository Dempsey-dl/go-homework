// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
✅  二分查找 (Binary Search)
题目描述：在一个有序数组中查找目标值。
*/

contract BinarySearch {

     function binary(uint256[] memory nums,uint256 target) public pure returns (int256) {
        uint256 left = 0;
        uint256 right = nums.length;

        while (left < right) {
            uint256 mid = left + (right - left) / 2; // 避免溢出
            
            if (nums[mid] == target) {
                return int256(mid); // 找到目标，返回索引
            } else if (nums[mid] < target) {
                left = mid + 1; // 目标在右半部分
            } else {
                right = mid; // 目标在左半部分
            }
        }
        
        return -1; // 未找到目标
     }

}