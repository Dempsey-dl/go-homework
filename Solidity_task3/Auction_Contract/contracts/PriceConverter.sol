// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorInterface.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

library PriceConcerter  {
    function covertToUSD(  
        uint amount,
        AggregatorV3Interface priceFeed
    ) internal view returns (uint256) {
        (, int256 price, , , ) = priceFeed.latestRoundData();
        uint8 decimals = priceFeed.decimals();
        return (amount * uint256(price)) / (10 ** decimals);
    }
}