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
      url: "https://sepolia.infura.io/v3/3cd6670397b442a3b321cdb81cb3074d",
      accounts: ["cc4d1984ff43567efeb9edb41f9ee1658c6f1016e7c62a98eba7e9738a83057c"],
      chainId: 11155111,
      gas: "auto",
      gasPrice: "auto"
    }
  }
};
