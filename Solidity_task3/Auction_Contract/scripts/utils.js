const fs = require("fs");
const path = require("path");

const deploymentFile = path.resolve(__dirname, "../deploy/.cache/deployment-info.json");

async function saveDeploymentInfo(info) {
  const existing = await getDeploymentInfo();
  fs.writeFileSync(deploymentFile, JSON.stringify({ ...existing, ...info }, null, 2));
}

async function getDeploymentInfo() {
  try {
    return JSON.parse(fs.readFileSync(deploymentFile, "utf8"));
  } catch (e) {
    return {};
  }
}

module.exports = { saveDeploymentInfo, getDeploymentInfo };