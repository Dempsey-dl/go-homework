**MetaNodeStake**
1.概览
支持本币  native token  erc20token   eth  抵押
采用可升级合约架构  UUPS + openzeppelin
动态调整区块奖励权重
提供延时解锁机制 （防止挤兑攻击）


2.技术栈:
solidity 0.8.20 （安全数学运算）
openzeppelin库  
UUPSUpgradeable  可升级代理模式
AccessControl    权限控制
SafaERC20        安全转账
Pausable         紧急暂停功能




说明  :
质押合约  起源是   DEX需要流动性  需要 LP（流动性提供商）给池子增加流动性 池子变深 交易的滑点就小     为了提升流动性  吸引LP    然后就做了流动性挖矿   DEX协议比较风行的时候    sushiswap开启  这时开始做流动性挖矿    



LP  提供流动性时    流动性池子返回给LP Token (提供流动性的权益凭证 可以长期持有赚取手续费  但也有无偿损失的风险)

平台开启了一个活动  把LP抵押给平台    然后你可以获取平台币


池子权重调整:  权重高的分的Token多



延时解锁机制:  当市场出现恐慌->大量用户同时unstake->质押资产被同时提取->导致流动性枯竭协议破产
延时解锁效果:  通过强制等待期分散提款压力,避免资金瞬时外流
