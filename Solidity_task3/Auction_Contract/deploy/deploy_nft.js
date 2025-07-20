const { ethers, getNamedAccounts, deployments } = require("hardhat")

module.exports = async ({ getNamedAccounts }) => {
    const { save } = deployments;

    const NFT = await ethers.getContractFactory("NFT");
    const nft = await NFT.deploy();
    await nft.waitForDeployment();
    const NFTAddress = await nft.getAddress();
    console.log("NFT 合约地址:", NFTAddress);


    // await save("NFTaddress", {
    //     abi: nft.interface.format("json"),
    //     address: NFTAddress
    // });
}