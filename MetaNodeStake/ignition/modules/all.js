const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");


module.exports = buildModule("MetaNodeModule", (m) => {
    const MetaNode = m.contract("MetaNode");
    return { MetaNode };
})