// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NFT is ERC721URIStorage {
    uint256 private tokenID;

    constructor() ERC721("NFT Collection", "NFT") {}

    function mint(address to, string memory tokenURI)  public   returns (uint256) {
        _safeMint(to, tokenID);
        _setTokenURI(tokenID, tokenURI);
        tokenID++;
        return tokenID - 1; // 返回实际铸造的tokenID
    }
} 