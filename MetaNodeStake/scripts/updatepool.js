const { ethers } = require("hardhat");

async function main() {

    const stake = await ethers.getContractAt("MetaNodeStake", "0x3Caaaa808c94f05E61074e0d416F7ae1d3D6aF69");

    const tx = await stake.updatePool(0, ethers.parseEther("0.000001"), 10);

    console.log(tx);

}

main();