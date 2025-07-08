// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract intToRoman {
    struct RomanNumeral {
        uint256 value;
        string symbol;
    }
    RomanNumeral[] private RomanNumerals; 
    constructor() {
        RomanNumerals.push(RomanNumeral(1000, "M"));
        RomanNumerals.push(RomanNumeral(900, "CM"));
        RomanNumerals.push(RomanNumeral(500, "D"));
        RomanNumerals.push(RomanNumeral(400, "CD"));
        RomanNumerals.push(RomanNumeral(100, "C"));
        RomanNumerals.push(RomanNumeral(90, "XC"));
        RomanNumerals.push(RomanNumeral(50, "L"));
        RomanNumerals.push(RomanNumeral(40, "XL"));
        RomanNumerals.push(RomanNumeral(10, "X"));
        RomanNumerals.push(RomanNumeral(9, "IX"));
        RomanNumerals.push(RomanNumeral(5, "V"));
        RomanNumerals.push(RomanNumeral(4, "IV"));
        RomanNumerals.push(RomanNumeral(1, "I"));  
    }

    function intToman(uint256 value) public view returns (string memory) {
        require(value > 1 && value < 4000, "Number must be between 1 and 3999");
        string memory Roman;

        for (uint i = 0 ; i < RomanNumerals.length; i++ )  {
            RomanNumeral memory cur = RomanNumerals[i];
            while (value >= cur.value)
            {
                Roman = string(abi.encodePacked(Roman,cur.symbol));   
                value -= cur.value;
            }
        }
        return Roman;
    }

}