// 合约初始化：验证合约部署后的初始状态

// 存款功能：测试正常存款和边界条件

// 奖励计算：验证奖励计算的准确性

// 解押流程：测试解押请求和实际提款

// 奖励领取：验证奖励领取功能

// 暂停机制：测试管理员的暂停功能

// 合约升级：验证可升级性


// 初始化   设置了   区块区间  每个区块产生的奖励   设置奖励的代币   包括权限的设置  升级 初始化

// 添加池子  代币地址  权重   最小抵押数量  解锁区块  是否更新池子

// 抵押   解压    获取奖励   提现   领取奖励  暂停机制   合约升级


const { ethers, upgrades, network } = require("hardhat");
const { expect } = require("chai");

describe("MetaNodeContract Test", function () {
    let LPToken;
    let lpToken;
    let MetaNodeStake;
    let metaNodeStake;
    let MetaNodeToken;
    let metaNodeToken;
    let owner, user1, user2;

    // 部署合约的初始参数
    const START_BLOCK = 100;
    const END_BLOCK = 1000;
    const REWARD_PER_BLOCK = ethers.parseEther("1"); // 1 token per block
    const MIN_DEPOSIT = ethers.parseEther("1");
    const LOCKED_BLOCKS = 50;
    const DEFAULT_WEIGHT = 100;

    before(async function () {
        [owner, user1, user2] = await ethers.getSigners();

        // 部署 LP  代币合约
        LPToken = await ethers.getContractFactory("LP");
        lpToken = await LPToken.deploy();
        await lpToken.waitForDeployment();

        await lpToken.mint(user1.address, ethers.parseEther("10"));
        await lpToken.mint(user2.address, ethers.parseEther("20"));

        // 部署奖励代币合约
        MetaNodeToken = await ethers.getContractFactory("MetaNode");
        metaNodeToken = await MetaNodeToken.deploy();
        metaNodeToken.waitForDeployment();
        // 部署Stake可升级合约
        MetaNodeStake = await ethers.getContractFactory("MetaNodeStake");
        metaNodeStake = await upgrades.deployProxy(MetaNodeStake, [
            metaNodeToken.target,
            START_BLOCK,
            END_BLOCK,
            REWARD_PER_BLOCK
        ]);

        const ver = await metaNodeStake.currentVerSion();
        console.log("当前版本", ver);

        await network.provider.send("hardhat_mine", ["0xc8"]);
        // 添加池子
        await metaNodeStake.addPool(ethers.ZeroAddress, DEFAULT_WEIGHT, MIN_DEPOSIT, LOCKED_BLOCKS, true);

        await network.provider.send("hardhat_mine", ["0x32"]);

        await metaNodeStake.addPool(await lpToken.getAddress(), DEFAULT_WEIGHT, MIN_DEPOSIT, LOCKED_BLOCKS, false);

        const startBlock = await ethers.provider.getBlockNumber();
        console.log("初始区块号:", startBlock); // 通常是 0

        const out = await metaNodeStake.getMultiplier(START_BLOCK, startBlock);
        console.log("挖矿:", out);


        // await metaNodeToken.transferFrom(owner.address, metaNodeStake.target, ethers.parseEther("900"));

        await metaNodeToken.transfer(metaNodeStake.target, ethers.parseEther("1000"));

    });




    describe("初始化测试", function () {
        it("应该正确初始化参数", async function () {
            expect(await metaNodeStake.getStartBlock()).to.equal(START_BLOCK);
            expect(await metaNodeStake.getendBlock()).to.equal(END_BLOCK);
            expect(await metaNodeStake.getOutput()).to.equal(REWARD_PER_BLOCK);
            expect(await metaNodeStake.poolLen()).to.equal(2);
        });
    })

    describe("抵押测试", function () {
        // it("ETH正确抵押",async function() {
        //    await metaNodeStake.depositETH();
        // });
        it("ERC正确抵押", async function () {
            await lpToken.connect(user1).approve(metaNodeStake.target, ethers.parseEther("4"));
            const amiunt = await lpToken.allowance(user1.address, metaNodeStake.target);
            await metaNodeStake.connect(user1).deposit(1, ethers.parseEther("3"));
            const stamount = await metaNodeStake.connect(user1).getStoke(1);
            expect(stamount).to.equal(ethers.parseEther("3"));
        });
    });

    describe("获取奖励", function () {
        it("应该正确获取奖励值", async function () {
            await network.provider.send("hardhat_mine", ["0x32"]);

            const pendingToken = await metaNodeStake.getPendingMetaNode(1, user1.address);
            console.log("获取的奖励", pendingToken);
        })
    })

    describe("解押", function () {
        it("正确解压", async function () {
            const balance = await lpToken.balanceOf(user1.address);
            console.log("当前余额", balance);

            await metaNodeStake.connect(user1).unStake(1, ethers.parseEther("1"));


            await metaNodeStake.connect(user1).withdraw(1);
            await network.provider.send("hardhat_mine", ["0x32"]);
            await metaNodeStake.connect(user1).withdraw(1);


            const balanceafter = await lpToken.balanceOf(user1.address);
            console.log("解压50个区块后 余额", balanceafter)
        })
    })

    describe("领取奖励", function () {
        it("领取奖励", async function () {
            // 查看当前有多少奖励
            const pendingToken = await metaNodeStake.getPendingMetaNode(1, user1.address);
            console.log("奖励余额", pendingToken);
            // 领取奖励 
            await metaNodeStake.connect(user1).cliam(1);

            const pendingToken1 = await metaNodeStake.getPendingMetaNode(1, user1.address);
            console.log("奖励后余额", pendingToken1);

            // 查看余额  
            const currentbalance = await metaNodeStake.getBalance(user1.address);
            console.log("当前余额", currentbalance);

            // expect(currentbalance).to.equal(pendingToken);
        })
    })


    describe("暂停", function () {
        it("提现暂停", async function () {
            await metaNodeStake.withdrawPause();
            await expect(
                metaNodeStake.connect(user1).withdraw(1)
            ).to.be.revertedWith("withdraw already paused");
        })
        it("领奖暂停", async function () {
            await metaNodeStake.cliamed();
            await expect(
                metaNodeStake.connect(user1).cliam(1)
            ).to.be.revertedWith("cliam already paused");
        })
    })

    describe("合约升级", function () {
        it("合约应该正常升级到V2", async function () {
            const MetaNodeStakeV2Factory = await ethers.getContractFactory("MetaNodeStakeV2");
            const metaNodeStakeV2 = await upgrades.upgradeProxy(metaNodeStake.target, MetaNodeStakeV2Factory);

            console.log("当前版本:",await metaNodeStakeV2.currentVerSion());
            // 验证升级后合约仍然正常工作
            expect(await metaNodeStakeV2.poolLength()).to.equal(2);
        })
    })
})


