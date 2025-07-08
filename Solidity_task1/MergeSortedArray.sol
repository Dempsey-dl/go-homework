// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
/*
✅  合并两个有序数组 (Merge Sorted Array)
题目描述：将两个有序数组合并为一个有序数组。
*/
contract MergeSortedArry {
    
    function mergearr(uint[] memory arr1,uint[] memory arr2) public pure returns (uint[] memory) {
        uint len1 = arr1.length;
        uint len2 = arr2.length;
        uint[] memory totalarr = new uint[](len1 + len2);
        uint i = 0;
        uint j = 0;
        uint k = 0;


        while (i < len1 && j < len2) {
            if (arr1[i] <= arr2[j]) {
                totalarr[k++] = arr1[i++];
            } else {
                totalarr[k++] = arr2[j++];
            }
        }


        while (i < len1) {
            totalarr[k++] = arr1[i++];
        }
         while (j < len2) {
             totalarr[k++] = arr2[j++];
        }

        return totalarr;
    }


}