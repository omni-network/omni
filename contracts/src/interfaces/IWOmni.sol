// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title IWOmni
 * @notice Interface for WOMNI, modifed interface of WETH9
 * @dev Changes: Renamed "WETH" to "WOmni"
 * @custom:attribution https://github.com/ethereum-optimism/optimism/blob/develop/packages/contracts-bedrock/src/dispute/interfaces/IWETH.sol
 */
interface IWOmni {
    /**
     * @notice Emitted when an approval is made.
     * @param src The address that approved the transfer.
     * @param guy The address that was approved to transfer.
     * @param wad The amount that was approved to transfer.
     */
    event Approval(address indexed src, address indexed guy, uint256 wad);

    /**
     * @notice Emitted when a transfer is made.
     * @param src The address that transferred the WOMNI.
     * @param dst The address that received the WOMNI.
     * @param wad The amount of WOMNI that was transferred.
     */
    event Transfer(address indexed src, address indexed dst, uint256 wad);

    /**
     * @notice Emitted when a deposit is made.
     * @param dst The address that deposited the WOMNI.
     * @param wad The amount of WOMNI that was deposited.
     */
    event Deposit(address indexed dst, uint256 wad);

    /**
     * @notice Emitted when a withdrawal is made.
     * @param src The address that withdrew the WOMNI.
     * @param wad The amount of WOMNI that was withdrawn.
     */
    event Withdrawal(address indexed src, uint256 wad);

    /**
     * @notice Returns the name of the token.
     * @return The name of the token.
     */
    function name() external view returns (string memory);

    /**
     * @notice Returns the symbol of the token.
     * @return The symbol of the token.
     */
    function symbol() external view returns (string memory);

    /**
     * @notice Returns the number of decimals the token uses.
     * @return The number of decimals the token uses.
     */
    function decimals() external pure returns (uint8);

    /**
     * @notice Returns the balance of the given address.
     * @param owner The address to query the balance of.
     * @return The balance of the given address.
     */
    function balanceOf(address owner) external view returns (uint256);

    /**
     * @notice Returns the amount of WOMNI that the spender can transfer on behalf of the owner.
     * @param owner The address that owns the WOMNI.
     * @param spender The address that is approved to transfer the WOMNI.
     * @return The amount of WOMNI that the spender can transfer on behalf of the owner.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @notice Allows WOMNI to be deposited by sending ether to the contract.
     */
    function deposit() external payable;

    /**
     * @notice Withdraws an amount of OMNI.
     * @param wad The amount of OMNI to withdraw.
     */
    function withdraw(uint256 wad) external;

    /**
     * @notice Returns the total supply of WOMNI.
     * @return The total supply of WOMNI.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @notice Approves the given address to transfer the WOMNI on behalf of the caller.
     * @param guy The address that is approved to transfer the WOMNI.
     * @param wad The amount that is approved to transfer.
     * @return True if the approval was successful.
     */
    function approve(address guy, uint256 wad) external returns (bool);

    /**
     * @notice Transfers the given amount of WOMNI to the given address.
     * @param dst The address to transfer the WOMNI to.
     * @param wad The amount of WOMNI to transfer.
     * @return True if the transfer was successful.
     */
    function transfer(address dst, uint256 wad) external returns (bool);

    /**
     * @notice Transfers the given amount of WOMNI from the given address to the given address.
     * @param src The address to transfer the WOMNI from.
     * @param dst The address to transfer the WOMNI to.
     * @param wad The amount of WOMNI to transfer.
     * @return True if the transfer was successful.
     */
    function transferFrom(address src, address dst, uint256 wad) external returns (bool);
}
