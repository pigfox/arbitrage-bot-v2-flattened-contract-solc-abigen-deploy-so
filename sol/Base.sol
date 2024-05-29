// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract Somemath{
    uint private q;
    constructor(){}

    event Adding(uint x, uint y);
    event SettingQ(uint q);
    function setQ(uint _q)external{
        emit SettingQ(_q);
        q = _q;
    }

    function getQ()external view returns (uint){
        return q;
    }

    function add(uint _x, uint _y) external returns(uint){
        emit Adding(_x, _y);
        return _x+_y;
    }
}

contract Stringer{
    string private x = "abc";
    constructor(){}
    function concat(string memory s)external view returns(string memory){
        return string(abi.encodePacked(x, s));
    }
}

contract Base is Somemath, Stringer{
    event Created(address owner, address delployedAt);
    constructor(){
        emit Created(msg.sender, address (this));
    }
}