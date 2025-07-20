// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/interfaces/IERC721.sol";
import "./PriceConverter.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract Auction is Initializable {
    event BidPlaced(
        address indexed bidder,
        uint256 amount,
        address currency,
        uint256 usdAmount
    );

    event AuctionEnded(address indexed winner, uint256 amount);

    using PriceConcerter for uint256;
    struct AuctionData {
        //购买者
        address seller;
        //起拍价
        uint256 startPrice;
        //开始时间
        uint256 startTime;
        //拍卖总时间
        uint256 duration;
        //结束时间
        uint256 endtime;
        //拍卖状态
        bool isend;
        //支付方式
        address paymentAddr;
        //最高出价者
        address highestbidder;
        //最高价
        uint256 highestbid;
        //所拍物品
        address nftAddr;
        //物品ID
        uint256 tokenID;
    }

    AggregatorV3Interface public priceFeet;

    AuctionData public auction;

    address public owner;

    function initialize() public initializer {
        owner = msg.sender;
    }

    function CreateAuction(
        address _seller,
        uint256 _startPrice,
        uint256 _duration,
        address _payment,
        address _nftaddr,
        uint256 _tokenID,
        address _priceFeet
    ) public virtual {
        auction = AuctionData({
            seller: _seller,
            startPrice: _startPrice,
            startTime: block.timestamp,
            duration: _duration,
            endtime: 0,
            isend: false,
            paymentAddr: _payment,
            highestbidder: address(0),
            highestbid: 0,
            nftAddr: _nftaddr,
            tokenID: _tokenID
        });
        priceFeet = AggregatorV3Interface(_priceFeet);
        IERC721(_nftaddr).transferFrom(_seller, address(this), _tokenID);
    }

    function Bid(uint256 amount) external payable virtual {
        // 是否在拍卖期内
        require(
            !auction.isend &&
                block.timestamp <= auction.startPrice + auction.duration,
            "exceed auction time"
        );

        uint256 bidamount;

        if (auction.paymentAddr == address(0)) {
            require(msg.value == amount, "ETH must be equl amount");
            bidamount = amount.covertToUSD(priceFeet);
        } else {
            require(msg.value == 0, "ETH not required for ERC20 bids");
            IERC20(auction.paymentAddr).transferFrom(
                msg.sender,
                address(this),
                amount
            );
            bidamount = amount.covertToUSD(priceFeet);
        }
        uint256 strtUSD = auction.startPrice.covertToUSD(priceFeet);
        uint256 hightUSD = auction.highestbid.covertToUSD(priceFeet);

        // 是否大于上一次拍卖
        require(bidamount > strtUSD, "Bid must be at least start price");
        require(
            bidamount > hightUSD,
            "Bid must be at least 5% higher than current bid"
        );
        // 退回
        if (auction.paymentAddr == address(0)) {
            //如果是ETH退回ETH
            payable(auction.highestbidder).transfer(auction.highestbid);
        } else {
            IERC20(auction.paymentAddr).transferFrom(
                address(this),
                auction.highestbidder,
                amount
            );
        }
        // 更新记录
        auction.highestbidder = msg.sender;
        auction.highestbid = amount;

        // 日志
        emit BidPlaced(msg.sender, amount, auction.paymentAddr, bidamount);
    }

    function endAuction() external virtual {
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
            if (auction.paymentAddr == address(0)) {
                // 给卖家ETH
                payable(auction.seller).transfer(auction.highestbid);
                //给卖家 ERC20
                IERC20(auction.paymentAddr).transferFrom(
                    address(this),
                    auction.seller,
                    auction.highestbid
                );
                emit AuctionEnded(auction.highestbidder, auction.highestbid);
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

    function Version() external pure virtual returns (string memory) {
        return "V1.0.0";
    }
}
