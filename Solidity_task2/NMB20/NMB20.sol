// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


contract NMB20 {
    //余额
    mapping (address => uint) _balances;
    //授权
    mapping (address => mapping (address => uint)) _allowances;
    
    uint private _totalSupply;
    string private _name;
    string private _symbol;
    address private owner;


    // 转账事件
    event Transfer(address indexed from, address indexed to, uint256 value);
    
    // 授权事件
    event Approval(address indexed owner, address indexed spender, uint256 value);

    constructor(string memory name_,string memory symbol_,uint initToken) {
        _name = name_;
        _symbol = symbol_;
        owner = msg.sender;
        _mint(owner,initToken);
    }

    function name() public view virtual  returns (string memory) {
        return _name;
    }

    function symbol() public view virtual  returns (string memory) {
        return _symbol;
    }

    function totalSupply() public view virtual  returns (uint) {
        return _totalSupply;
    }
    
     function decimals() public view virtual returns (uint8) {
        return 18;
    }

    function balance(address addr) public view virtual returns(uint) {
        return _balances[addr];
    }

    function transfer(address to,uint amount) public returns (bool) {
        require(to != address(0),"transfer to the zero address");
        require(_balances[msg.sender] >= amount,"owner amount to exceeds balance");

        _balances[msg.sender] -= amount;
        _balances[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function approve(address to,uint amount) public returns (bool) {
        require(to != address(0),"address cannot zero");
        require(_balances[msg.sender] > 0,"balances > 0");
        require(amount > 0,"amount > 0");

        _allowances[owner][to] = amount;

        emit Approval(owner, to, amount);

        return true;
    } 

    function allowance(address to) public view returns (uint) {
        return _allowances[msg.sender][to];
    }


    function transferFrom(address from,address to,uint amount)public  returns (bool) {
        require(_balances[from] >= amount,"0");
        require(to != address(0),"0");
        require(_allowances[from][msg.sender] >= amount,"allowances >= 0");

        _balances[from] -= amount;
        _balances[to] += amount;
        _allowances[from][msg.sender] -= amount;

        emit Transfer(from, to, amount);

        return true;
    }


    function mint(address user,uint amount) public {
        require(owner == user,"address is zero");
        _mint(user, amount);
    }

     // 内部铸币函数
    function _mint(address user,uint amount) internal  {
        require(amount != 0,"amount cannot is empty");
        _totalSupply += amount;
        _balances[user] += amount;
        emit Transfer(address(0), user, amount);
    }

}
