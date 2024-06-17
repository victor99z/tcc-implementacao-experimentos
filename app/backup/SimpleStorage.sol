// SPDX-License-Identifier: GPL-3.0

pragma solidity >0.7.0 < 0.9.0;

contract SimpleStorage {
  uint public storedData;

  constructor(uint initVal) {
    storedData = initVal;
  }

  function set(uint x) public {
    storedData = x;
  }

  function get() view public returns (uint) {
    return storedData;
  }
}