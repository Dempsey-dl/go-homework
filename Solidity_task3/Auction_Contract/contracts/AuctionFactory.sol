// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;  // 使用最新稳定版

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "./Auction.sol";
 
contract AuctionFactory {
    address public auctionImplementation;
    address public proxyAdmin;
    
    mapping(address => mapping(uint256 => address)) public auctions;
    address[] public allAuctions;

    event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address indexed seller, address auctionAddress);

    constructor(address _auctionImplementation, address _proxyAdmin) {
        auctionImplementation = _auctionImplementation;
        proxyAdmin = _proxyAdmin;
    }

    function createAuction(
        uint256 _startPrice,
        uint256 _duration,
        address _payment,
        address _nftaddr,
        uint256 _tokenID,
        address _priceFeet
    ) external returns (address) {
        require(auctions[_nftaddr][_tokenID] == address(0), "Auction exists");
         
        IERC721(_nftaddr).transferFrom(msg.sender, address(this), _tokenID);

        bytes memory data = abi.encodeWithSelector(
            Auction.initialize.selector,
            msg.sender,
            _startPrice,
            _duration,
            _payment,
            _nftaddr,
            _tokenID, 
            _priceFeet
        );

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            auctionImplementation,
            proxyAdmin,
            data
        );

        address auctionAddress = address(proxy);
        auctions[_nftaddr][_tokenID] = auctionAddress;
        allAuctions.push(auctionAddress);
 
        emit AuctionCreated(_nftaddr, _tokenID, msg.sender, auctionAddress);
        return auctionAddress;
    }
 

    function getAllAuctions() external view returns (address[] memory) {
        return allAuctions;
    }
    
    function updateImplementation(address _newImplementation) external {
        // require(msg.sender == proxyAdmin, "Only ProxyAdmin can update");
        auctionImplementation = _newImplementation;
    }

}  