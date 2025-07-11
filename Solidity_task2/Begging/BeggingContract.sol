// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
// 创建一个名为 BeggingContract 的合约。
// 合约应包含以下功能：
// 一个 mapping 来记录每个捐赠者的捐赠金额。
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。


 
contract BeggingContract {
  
    event Donation(address indexed donor,uint amount);

    // 捐款活动 宗旨
    string private  _reason;

    // 最大捐款数
    uint private _maxAmount;
    
    // 捐款总额
    uint private _currentDonate;

    // 捐款信息
    mapping (address => uint) private _donation;

    address private owner;

    struct TopDontor {
        address dontor;
        uint amount;
    }
    //前三
    TopDontor[3] private topDontors;

    // 额外挑战3：时间限制
    uint256 private  donationStartTime;
    uint256 private  constant DONATION_PERIOD = 7 days;


    //捐款理由 和  筹集金额
    constructor(string memory reason_,uint maxAmount_) {
        _reason = reason_;
        _maxAmount = maxAmount_;
        owner = msg.sender;
        donationStartTime = block.timestamp;
    }
    
    modifier onlyOwner() {
        require(msg.sender == owner,"only owenr can call this function");
        _;
    }

    modifier donationPeriod() {
        require(block.timestamp <= donationStartTime + DONATION_PERIOD,"Donation period has end");
        _;
    }

    function donationReason() public view virtual returns (string memory) {
        return _reason;
    }

    function maxAmount() public view virtual returns (uint) {
        return _maxAmount;
    }

    function currentDonate() public view virtual returns (uint) {
        return _currentDonate;
    }

    function getDonation(address donator) public view virtual returns (uint) {
        return _donation[donator];
    }

    function getadderss() public view returns (address) {
        return msg.sender;
    }

    //更新排名
    function updateTopDonors(address dontor, uint amount) private {
        
        for (uint i = 0; i < 3; i++) {
          if (amount > topDontors[i].amount)
          {
              for (uint j = 2; j > i; j--) {
                topDontors[j] = topDontors[j - 1];
              }

              topDontors[i] = TopDontor(dontor,amount);
              break;
          }
        }
    }

    //捐款
    function donate() public payable donationPeriod {
        require(msg.value > 0,"The donate amount must be greeter than zero");

        _donation[msg.sender] += msg.value;
        _currentDonate += msg.value;

        //排名
        updateTopDonors(msg.sender,msg.value);

        emit Donation(msg.sender,msg.value);
    }

    function withdraw() public payable onlyOwner {

        uint balance = address(this).balance;
        require(balance > 0,"The amount cannot must be greeter than balance");

        _currentDonate -= balance;

        payable(owner).transfer(balance);
    }

    function getPeriod() public view returns (uint) {
        if (block.timestamp >= donationStartTime + DONATION_PERIOD) {
            return 0;
        }
        return donationStartTime + DONATION_PERIOD - block.timestamp;
    }
}



