const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("NFT Auction System", function() {
  let NFT, nft;
  let Auction, auction, auctionV2;
  let AuctionFactory, factory;
  let ProxyAdmin, proxyAdmin;
  
  let owner, seller, bidder1, bidder2;
  let MockAggregator, mockAggregator;
  let MockERC20, mockERC20;

  before(async function() {
    [owner, seller, bidder1, bidder2] = await ethers.getSigners();
    
    // 部署NFT合约
    NFT = await ethers.getContractFactory("NFT");
    nft = await NFT.deploy();
    
    // 部署价格聚合器mock
    MockAggregator = await ethers.getContractFactory("MockAggregator");
    mockAggregator = await MockAggregator.deploy(8, 200000000); // 8 decimals, $200
    
    // 部署ERC20 mock
    MockERC20 = await ethers.getContractFactory("MockERC20");
    mockERC20 = await MockERC20.deploy("Mock Token", "MOCK", 18);
    
    // 部署拍卖合约V1
    Auction = await ethers.getContractFactory("Auction");
    auction = await upgrades.deployProxy(Auction, [], {initializer: 'initialize'});
    
    // 部署拍卖合约V2
    AuctionV2 = await ethers.getContractFactory("AuctionV2");
    auctionV2 = await upgrades.upgradeProxy(auction.address, AuctionV2);
    
    // 部署ProxyAdmin
    ProxyAdmin = await ethers.getContractFactory("ProxyAdmin");
    proxyAdmin = await ProxyAdmin.deploy();
    
    // 部署工厂合约
    AuctionFactory = await ethers.getContractFactory("AuctionFactory");
    factory = await AuctionFactory.deploy(auctionV2.address, proxyAdmin.address);
    
    // 给测试用户分配一些ERC20代币
    await mockERC20.mint(bidder1.address, ethers.utils.parseEther("1000"));
    await mockERC20.mint(bidder2.address, ethers.utils.parseEther("1000"));
  });
});