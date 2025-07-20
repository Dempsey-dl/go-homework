const { ethers, upgrades } = require("hardhat");
const { getDeploymentInfo, saveDeploymentInfo } = require("../scripts/utils");

module.exports = async function ({ getNamedAccounts }) {
  const { deployer } = await getNamedAccounts();
  const { ProxyAddress, ImpAddress } = await getDeploymentInfo();

  console.log("\n========== 升级实现合约阶段 ==========");
  console.log("代理地址:", ProxyAddress, "实现地址:", ImpAddress)

  // 1. 部署新版本 AuctionV2
  const AuctionV2 = await ethers.getContractFactory("AuctionV2");
  const auctionV2 = await upgrades.upgradeProxy(ProxyAddress, AuctionV2);
  await auctionV2.waitForDeployment();
  console.log("Auction 实现合约 V2 已部署:", await auctionV2.getAddress());
  const auctionV2Adress = await upgrades.erc1967.getImplementationAddress(ProxyAddress)

  // 3. 验证升级
  console.log("新实现合约地址:", auctionV2Adress);


  // 更新部署信息
  await saveDeploymentInfo({
    ProxyAddress: ProxyAddress,
    ImpAddress: ImpAddress,
    auctionV2Adress: auctionV2Adress
  });
    console.log("升级完成！");

};

module.exports.tags = ["upgrade"];