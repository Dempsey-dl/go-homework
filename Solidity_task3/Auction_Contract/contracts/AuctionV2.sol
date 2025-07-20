// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "./Auction.sol";


contract AuctionV2 is Auction {
    uint256 public feePercentage;

    function setFeePercentage(uint256 _feePercentage) external {
        require(_feePercentage <= 100, "Fee cannot exceed 100%");
        feePercentage = _feePercentage;
    }

    function endAuction() external override {
        // 判断是否在时间内
        require(
            block.timestamp >= auction.startPrice + auction.duration,
            "Auction not end yet"
        );
        // 判断合约状态
        require(!auction.isend, "Auction areadly end");
        // 结束时间和 开始甲duration
        auction.endtime = block.timestamp;
        // 设置结束状态
        auction.isend = true;
        // 判断又没有拍卖纪录
        if (auction.highestbidder != address(0)) {
            // 有的话  把nft给最高者  把钱给卖家
            IERC721(auction.nftAddr).transferFrom(
                address(this),
                auction.highestbidder,
                auction.tokenID
            );
            // 扣除手续费
            uint256 feeAmount = (auction.highestbid * feePercentage) / 100;
            uint256 sellerAmount = auction.highestbid - feeAmount;

            if (auction.paymentAddr == address(0)) {
                // 给卖家ETH
                payable(auction.seller).transfer(sellerAmount);
                payable(owner).transfer(feeAmount);
                emit AuctionEnded(auction.highestbidder, auction.highestbid);
            } else {
                //给卖家 ERC20
                IERC20(auction.paymentAddr).transferFrom(
                    address(this),
                    auction.seller,
                    auction.highestbid
                );

                IERC20(auction.paymentAddr).transferFrom(
                    address(this),
                    owner,
                    auction.highestbid
                );
            }
        } else {
            // 没有 nft原路返回
            IERC721(auction.nftAddr).transferFrom(
                address(this),
                auction.seller,
                auction.tokenID
            );
            emit AuctionEnded(address(0), 0);
        }
    }
    
    function Version() external pure override returns(string memory) {
        return "V2.0.0";
    }
}
