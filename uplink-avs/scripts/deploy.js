const hre = require("hardhat");

async function main() {
  const UpLinkAVS = await hre.ethers.getContractFactory("UpLinkAVS");
  const uplinkAVS = await UpLinkAVS.deploy();

  await uplinkAVS.deployed();
  console.log("UpLinkAVS deployed to:", uplinkAVS.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
