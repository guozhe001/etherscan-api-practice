pragma solidity ^0.6.0;

import "./ERC20Permit.sol";
import "./token/ERC20/ERC20Burnable.sol";
import "./access/Ownable.sol";

contract OneInch is ERC20Permit, ERC20Burnable, Ownable {
    constructor(address _owner) public ERC20("1INCH Token", "1INCH") EIP712("1INCH Token", "1") {
        _mint(_owner, 1.5e9 ether);
        transferOwnership(_owner);
    }

    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
}