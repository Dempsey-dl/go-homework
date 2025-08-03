const { ethers, upgrades } = require("hardhat")

async function main() {


    const WEIGHT = 2;
    const MIN = ethers.parseEther("0.000001");
    const UNLOCKBLOCK = 50;
    const ISUPDATE = true;

    const MetaNodeStake = await ethers.getContractAt("MetaNodeStake", "0x2EAB43DEFCa639056644Eb2E0bcCFCdB49f3b2b3");

    const pool = await MetaNodeStake.addPool(ethers.ZeroAddress, WEIGHT, MIN, UNLOCKBLOCK, ISUPDATE);
    console.log(pool);
}

main();