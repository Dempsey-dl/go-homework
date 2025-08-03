const { ethers, upgrades } = require("hardhat")



async function main() {
    const MetaNodeStake = await ethers.getContractAt("MetaNodeStake", "0x3Caaaa808c94f05E61074e0d416F7ae1d3D6aF69");
    const proxyAddress = await MetaNodeStake.getAddress();
    console.log("代理地址:", proxyAddress);

    const MetaNodeStakeV2 = await ethers.getContractFactory("MetaNodeStakeV2");
    const StakeV2 = await upgrades.upgradeProxy(proxyAddress, MetaNodeStakeV2);
    await StakeV2.waitForDeployment();


    console.log("当前地址:", await StakeV2.currentVerSion());
}

main();