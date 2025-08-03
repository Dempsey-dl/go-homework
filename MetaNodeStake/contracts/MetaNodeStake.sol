// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";


contract MetaNodeStake1 is
    Initializable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable
{
    using SafeERC20 for IERC20;
    using Math for uint256;
    using Address for address;

    // 角色设置
    bytes32 private ADMIN_ROLE;
    bytes32 private UPGRADE_ROLE;

    uint256 private ETH_PID;

    function currentVerSion() public pure virtual returns (string memory) {
        return "V1.0";
    }

    // Pool

    struct Pool {
        // 抵押代币地址
        address stTokenAddress;
        // 权重
        uint256 weight;
        // 最后更新区块
        uint256 lastRewardBlock;
        // 最小抵押值
        uint256 minStAmount;
        // 每个抵押代币获取的奖励
        uint256 accMetaNodePerSt;
        // 抵押代币总量
        uint256 stSupply;
        // 解压区块
        uint256 unStakeLockBlock;
    }
    // User
    struct User {
        // 代币数量
        uint256 stAmount;
        // 已完成的结算奖励
        uint256 finishedMetaNode;
        // 待领取的奖励
        uint256 pendingMetaNode;
        // 解压请求
        UnStakeRequestd[] request;
    }
    // UnStakeRequestd
    struct UnStakeRequestd {
        // 代币数
        uint256 amount;
        // 解锁区块
        uint256 unlockblock;
    }

    // 开始区块
    uint256 private _startblock;
    // 结束区块
    uint256 private _endblock;
    // 每个区块产出的MetaNode
    uint256 private _metaNodePerBlock;
    // 总权重
    uint256 private _totalWeight;
    // 体现暂停状态
    bool private _isWithdrawed;
    // 领取奖励状态
    bool private _isCliamed;
    // IERC20
    IERC20 private _MetaNode;
    // 池子信息
    Pool[] private _pool;
    // 用户信息
    mapping(uint256 => mapping(address => User)) _user;
    // 修饰符
    // checkpid

    modifier checkpid(uint256 pid) {
        require(pid < _pool.length, "pid invalid");
        _;
    }

    // whenNotWithdrawPaused
    modifier whenNotWithdrawPaused() {
        require(!_isWithdrawed, "withdraw already paused");
        _;
    }
    // whenNotCliamPaused
    modifier whenNotCliamPaused() {
        require(!_isCliamed, "cliam already paused");
        _;
    }

    // 事件
    event withdrawEvent(uint256 pid, address user, uint256 amount);

    // 初始化
    function initialize(
        IERC20 erc,
        uint256 start,
        uint256 end,
        uint256 tokenperblock
    ) public initializer {
        // 检查参数
        require(start <= end && tokenperblock > 0, "invalid paremters");
        // 角色权限初始化
        __AccessControl_init();
        // 可升级初始化
        __UUPSUpgradeable_init();
        // 设置 区间快

        ETH_PID = 0;
        ADMIN_ROLE = keccak256("admin_role");
        UPGRADE_ROLE = keccak256("upgrade_role");

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(UPGRADE_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);

        // 设置 每个区块产生多少奖励
        _startblock = start;
        _endblock = end;
        _metaNodePerBlock = tokenperblock;
        // 设置代币
        setMetaNode(erc);
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyRole(UPGRADE_ROLE) {}

    function setMetaNode(IERC20 MetaNode_) public onlyRole(ADMIN_ROLE) {
        _MetaNode = MetaNode_;
        // emit SetMetaNode(MetaNode);
    }

    function withdrawPause() public onlyRole(ADMIN_ROLE) {
        require(!_isWithdrawed, "withdraw already pasuse");
        _isWithdrawed = true;
    }

    function unWithdrawPause() public onlyRole(ADMIN_ROLE) {
        require(_isWithdrawed, "withdraw already unpasuse");
        _isWithdrawed = false;
    }

    function cliamed() public onlyRole(ADMIN_ROLE) {
        require(!_isCliamed, "withdraw already pasuse");
        _isCliamed = true;
    }

    function unWcliamed() public onlyRole(ADMIN_ROLE) {
        require(_isCliamed, "withdraw already unpasuse");
        _isCliamed = false;
    }

    function setStartBlock(uint256 start) public onlyRole(ADMIN_ROLE) {
        require(start < _endblock, "Invalid paremter");
        _startblock = start;
    }

    function setendBlock(uint256 end) public onlyRole(ADMIN_ROLE) {
        require(end > _startblock, "Invalid paremter");
        _endblock = end;
    }

    function setoOutput(uint256 output) public onlyRole(ADMIN_ROLE) {
        require(output > 0, "Invalid paremter");
        _metaNodePerBlock = output;
    }

    function getStartBlock()
        public
        view
        onlyRole(ADMIN_ROLE)
        returns (uint256)
    {
        return _startblock;
    }

    function getendBlock() public view onlyRole(ADMIN_ROLE) returns (uint256) {
        return _endblock;
    }

    function getOutput() public view onlyRole(ADMIN_ROLE) returns (uint256) {
        return _metaNodePerBlock;
    }

    function poolLen() public view onlyRole(ADMIN_ROLE) returns (uint256) {
        return _pool.length;
    }

    // 添加池子
    function addPool(
        address stAddress,
        uint256 weight,
        uint256 minStamount,
        uint256 unlockblock,
        bool isUpdate
    ) public onlyRole(ADMIN_ROLE) {
        // 检查参数
        if (_pool.length > 0) {
            require(stAddress != address(0), "staddress cannot equl zero");
        } else {
            require(stAddress == address(0), "staddress must be equl zero");
        }
        require(unlockblock > 0, "Invalid paremters");
        require(block.number < _endblock, "already end");
        // 更新全部区块

        // massUpdateBlock
        if (isUpdate) {
            massUpdateBlock();
        }

        uint256 updatblock = block.number > _startblock
            ? block.number
            : _startblock;
        _totalWeight = _totalWeight + weight;
        _pool.push(
            Pool({
                stTokenAddress: stAddress,
                weight: weight,
                lastRewardBlock: updatblock,
                minStAmount: minStamount,
                accMetaNodePerSt: 0,
                stSupply: 0,
                unStakeLockBlock: unlockblock
            })
        );
    }

    function getBalance(address user) public view returns (uint256) {
        return _MetaNode.balanceOf(user);
    }

    // 更新全部区块
    function massUpdateBlock() private onlyRole(ADMIN_ROLE) {
        uint256 len = _pool.length;
        for (uint256 pid = 0; pid < len; pid++) {
            updatePool(pid);
        }
    }

    function poolLength() public view returns (uint256) {
        return _pool.length;
    }

    // 更新池子 最小抵押  和  解压区块高度
    function updatePool(
        uint256 _pid,
        uint256 minSt,
        uint256 unlockblock
    ) public onlyRole(ADMIN_ROLE) {
        _pool[_pid].minStAmount = minSt;
        _pool[_pid].unStakeLockBlock = unlockblock;
    }

    function updatePool(uint256 _pid) public checkpid(_pid) {
        Pool storage pool = _pool[_pid];
        // 更新每个stoke能获得多少奖励
        if (block.number <= pool.lastRewardBlock) {
            return;
        }

        // get奖励(上次更新的区块 到现在的区块)
        uint256 metaNodeAmount_ = getMultiplier(
            pool.lastRewardBlock,
            block.number
        );

        // 按权重比 分配奖励 (当前池子的奖励)
        uint256 metaNodeAmountforPool = ((metaNodeAmount_ * pool.weight) /
            _totalWeight) * (1 ether);
        // 池子的奖励数量 除以 池子总抵押数量
        if (pool.stSupply > 0) {
            pool.accMetaNodePerSt =
                pool.accMetaNodePerSt +
                metaNodeAmountforPool /
                pool.stSupply;
        }
        // 更新区块
        pool.lastRewardBlock = block.number;
    }

    // 更新权重
    function updateWeight(
        uint256 _pid,
        uint256 weight
    ) public onlyRole(ADMIN_ROLE) {
        Pool storage pool = _pool[_pid];
        _totalWeight = _totalWeight - pool.weight + weight;
        pool.weight = weight;
    }

    // getMultiplier   区块区间 产生代币的数量
    function getMultiplier(
        uint256 from,
        uint256 to
    ) public view returns (uint256) {
        require(from < to, "Invalid paremters");
        if (from < _startblock) {
            from = _startblock;
        }
        if (to > _endblock) {
            to = _endblock;
        }
        require(from < to, "from must be greater than to");

        return (to - from) * _metaNodePerBlock;
    }

    // getPendingMetaNode  获取待领取奖励值
    function getPendingMetaNode(
        uint256 pid_,
        address user_
    ) public view checkpid(pid_) returns (uint256) {
        return getPendingMetaNodeByBlock(pid_, user_, block.number);
    }

    // getPendingMetaNodeByBlock
    function getPendingMetaNodeByBlock(
        uint256 pid_,
        address user_,
        uint256 block_
    ) public view checkpid(pid_) returns (uint256) {
        Pool storage pool = _pool[pid_];
        User storage user = _user[pid_][user_];
        uint256 accMetaNodePerST = pool.accMetaNodePerSt;
        if (pool.lastRewardBlock < block_ && pool.stSupply != 0) {
            // 先计算区间奖励  乘以权重比  除以 总抵押币
            uint256 multiplier = getMultiplier(pool.lastRewardBlock, block_);
            uint256 metaNodeAmount = (multiplier * pool.weight) / _totalWeight;
            accMetaNodePerST =
                accMetaNodePerST +
                (metaNodeAmount * (1 ether)) /
                pool.stSupply;
        }
        return
            (user.stAmount * accMetaNodePerST) /
            (1 ether) -
            user.finishedMetaNode +
            user.pendingMetaNode;
    }

    // 获取用户的质押数量
    function getStoke(
        uint256 pid_
    ) public view checkpid(pid_) returns (uint256) {
        return _user[pid_][msg.sender].stAmount;
    }

    // 抵押ETH
    function depositETH() public payable whenNotPaused {
        Pool memory pool = _pool[ETH_PID];
        require(
            pool.stTokenAddress == address(0),
            "Invalid staking token address"
        );
        uint256 amount = msg.value;
        require(amount >= pool.minStAmount, "deposit amount is too small");
        _deposit(ETH_PID, amount);
    }

    // 抵押代币
    function deposit(
        uint256 pid_,
        uint256 amount
    ) public whenNotPaused checkpid(pid_) {
        Pool memory pool = _pool[pid_];
        require(
            pool.stTokenAddress != address(0),
            "Invalid staking token address"
        );

        require(amount > pool.minStAmount, "deposit amount is too small");

        IERC20(pool.stTokenAddress).transferFrom(
            msg.sender,
            address(this),
            amount
        );

        _deposit(pid_, amount);
    }

    // _deposit
    function _deposit(uint256 pid_, uint256 amount) internal {
        Pool storage pool = _pool[pid_];
        User storage user = _user[pid_][msg.sender];

        // 更新池子
        updatePool(pid_);
        // 更新待领取奖励
        if (user.stAmount > 0) {
            uint256 pendingMetaNode_ = (user.stAmount * pool.accMetaNodePerSt) /
                (1 ether) -
                user.finishedMetaNode;
            if (pendingMetaNode_ > 0) {
                user.pendingMetaNode = pendingMetaNode_ + user.pendingMetaNode;
                // _MetaNode._mint(DEFAULT_ADMIN_ROLE,user.pendingMetaNode);
            }
        }
        user.stAmount = user.stAmount + amount;
        pool.stSupply = pool.stSupply + amount;
        // 更新已经结算奖励
        user.finishedMetaNode =
            (user.stAmount * pool.accMetaNodePerSt) /
            (1 ether);
    }

    // 解押
    function unStake(
        uint256 pid_,
        uint256 amount
    ) public whenNotPaused checkpid(pid_) whenNotWithdrawPaused {
        Pool storage pool = _pool[pid_];
        User storage user = _user[pid_][msg.sender];
        require(user.stAmount > amount, "Not enough balance");
        // 更新池子
        updatePool(pid_);

        uint256 pendingMetaNode_ = (user.stAmount * pool.accMetaNodePerSt) /
            (1 ether) -
            user.finishedMetaNode;
        if (pendingMetaNode_ > 0) {
            user.pendingMetaNode = pendingMetaNode_ + user.pendingMetaNode;
        }

        user.request.push(
            UnStakeRequestd({
                amount: amount,
                unlockblock: block.number + pool.unStakeLockBlock
            })
        );

        user.stAmount = user.stAmount - amount;
        user.finishedMetaNode =
            (user.stAmount * pool.accMetaNodePerSt) /
            (1 ether);
        pool.stSupply = pool.stSupply - amount;
    }

    // 提现
    function withdraw(
        uint256 pid_
    )
        public
        whenNotPaused
        checkpid(pid_)
        whenNotWithdrawPaused
        returns (uint256)
    {
        Pool storage pool = _pool[pid_];
        User storage user = _user[pid_][msg.sender];
        uint256 pidNum;
        uint256 withdrawvalue;
        for (uint i = 0; i < user.request.length; i++) {
            if (user.request[i].unlockblock > block.number) {
                break;
            }
            withdrawvalue = withdrawvalue + user.request[i].amount;
            pidNum++;
        }

        for (uint i = 0; i < user.request.length - pidNum; i++) {
            user.request[i] = user.request[i + pidNum];
        }

        for (uint i = 0; i < pidNum; i++) {
            user.request.pop;
        }

        if (withdrawvalue > 0) {
            if (pool.stTokenAddress == address(0)) {
                _safeETHtransfer(msg.sender, withdrawvalue);
            } else {
                IERC20(pool.stTokenAddress).safeTransfer(
                    msg.sender,
                    withdrawvalue
                );
            }
        }

        emit withdrawEvent(pid_, msg.sender, withdrawvalue);

        return withdrawvalue;
    }

    // 领取奖励
    function cliam(
        uint256 pid_
    ) public whenNotPaused checkpid(pid_) whenNotCliamPaused {
        Pool storage pool = _pool[pid_];
        User storage user = _user[pid_][msg.sender];

        // 更新池子
        updatePool(pid_);

        uint256 pendingmetnode = (user.stAmount * pool.accMetaNodePerSt) /
            (1 ether) -
            user.finishedMetaNode +
            user.pendingMetaNode;
        if (pendingmetnode > 0) {
            user.pendingMetaNode = 0;
            _safeTransfer(msg.sender, pendingmetnode);
        }

        user.finishedMetaNode =
            (user.stAmount * pool.accMetaNodePerSt) /
            (1 ether);
    }

    function _safeTransfer(address to, uint256 amount) internal {
        uint256 metaNodeAmount = _MetaNode.balanceOf(address(this));
        require(metaNodeAmount > 0, "No tokens to transfer");

        uint256 transferAmount = amount > metaNodeAmount
            ? metaNodeAmount
            : amount;
        require(_MetaNode.transfer(to, transferAmount), "Transfer failed"); // 强制检查返回值
    }

    // 安全转账 ETH
    function _safeETHtransfer(address to, uint256 amount) internal {
        (bool success, bytes memory data) = address(to).call{value: amount}("");
        require(success, "ETH transfer call failed");
        if (data.length > 0) {
            require(
                abi.decode(data, (bool)),
                "ETH transfer operation did not succeed"
            );
        }
    }
    // 安全转账 代币
}
