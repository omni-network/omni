// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

interface IOmniPortalAdmin {
    function admin() external view returns (address);
    function setAdmin(address admin) external;
    function fee() external view returns (uint256);
    function setFee(uint256 fee) external;
}
