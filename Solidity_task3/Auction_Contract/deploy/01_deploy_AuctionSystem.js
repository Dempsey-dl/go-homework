const { ethers, upgrades } = require("hardhat");
const { saveDeploymentInfo } = require("../scripts/utils");

module.exports = async function ({ getNamedAccounts, deployments }) {

  console.log("\n========== 部署合约阶段 ==========");

  const { deployer } = await getNamedAccounts();

  const Auction = await ethers.getContractFactory("Auction");
  const AuctionProxy = await upgrades.deployProxy(Auction, [], {
    initializer: "initialize"
  });
  const ProxyAddress = await AuctionProxy.getAddress();
  const ImpAddress = await upgrades.erc1967.getImplementationAddress(ProxyAddress);
  console.log("代理地址:", ProxyAddress, "实现地址:", ImpAddress)

  // 保存部署信息
  await saveDeploymentInfo({
    ProxyAddress: ProxyAddress,
    ImpAddress: ImpAddress,
  });
};

module.exports.tags = ["deploy"];