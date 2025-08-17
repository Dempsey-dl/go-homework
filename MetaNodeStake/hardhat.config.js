require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require("dotenv").config;
/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  namedAccounts: {
    deployer: 0,
    user1: 1,
    user2: 2
  },
  networks: {
    Sepolia: {
      // url: `https://sepolia.infura.io/v3/${process.env.INFURA}`,
      // accounts:[process.env.KEY]
      url: "https://sepolia.infura.io/v3/*",
      accounts: ["*"],
      chainId: 11155111,
      gas: "auto",
      gasPrice: "auto"
    }
  }
};
