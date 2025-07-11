// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract myNTF is  ERC721URIStorage {

    using Counters for Counters.Counter;
    Counters.Counter private _tokenids;
    constructor() ERC721("LemonNTF","NMB") { } 


    function mint(address recipient,string memory tokenUrl) public returns (address, uint, string memory) {
        _tokenids.increment();

        uint  newtokenid = _tokenids.current();
        _safeMint(recipient, newtokenid);
        _setTokenURI(newtokenid, tokenUrl);

        return (recipient,newtokenid,tokenUrl); 
    }   
}