const { ethers, upgrades } = require("hardhat");

async function main() {

    const [signer] = await ethers.getSigners();
    // 部署奖励代币    
    const MetaNodeToken = await ethers.getContractFactory("MetaNode");
    const metaNodeToken = await MetaNodeToken.deploy();
    await metaNodeToken.waitForDeployment();

    const metaNodeTokenAddress = await metaNodeToken.getAddress();

    console.log("奖励代币地址:", metaNodeTokenAddress);

    const STARTBLOCK = 8896088;
    const ENDBLOCK = 8900000;
    const tokenPerBlock = ethers.parseEther("1", 18);
    const MetaNodeStake = await ethers.getContractFactory("MetaNodeStake");
    const metaNodeStake = await upgrades.deployProxy(MetaNodeStake, [metaNodeTokenAddress, STARTBLOCK, ENDBLOCK, tokenPerBlock], {
        initializer: "initialize"
    });
    await metaNodeStake.waitForDeployment();
    const proxyAddress = await metaNodeStake.getAddress();

    console.log("代理地址:", proxyAddress);
    const implementationAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
    console.log("V1 实现合约地址:", implementationAddress);

    const balance = await metaNodeToken.balanceOf(signer.address);
    const tx = await metaNodeToken.transfer(metaNodeTokenAddress, balance)
    await tx.wait();
}

main().then(() => process.exit(0)).catch((error) => {
    console.error(error);
    process.exit(1);
});