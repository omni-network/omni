// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

interface IEIP712 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}

interface IMultiCall3 {
    struct Call {
        address target;
        bytes callData;
    }

    struct Result {
        bool success;
        bytes returnData;
    }

    function aggregate(Call[] calldata calls) external payable returns (uint256 blockNumber, bytes[] memory returnData);
    function tryAggregate(bool requireSuccess, Call[] calldata calls)
        external
        payable
        returns (Result[] memory returnData);

    function getBlockNumber() external view returns (uint256 blockNumber);
    function getLastBlockHash() external view returns (bytes32 blockHash);
    function getChainId() external view returns (uint256 chainid);
}

interface ICreate2Deployer {
    function deploy(uint256 value, bytes32 salt, bytes memory code) external;
    function computeAddress(bytes32 salt, bytes32 codeHash) external view returns (address);
}

interface ICreateX {
    function deployCreate2(bytes memory initCode) external payable returns (address newContract);
    function deployCreate2(bytes32 salt, bytes memory initCode) external payable returns (address newContract);
    function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) external view returns (address computedAddress);
}

interface ISafe_v130 {
    function getChainId() external view returns (uint256);
}

interface ISafeL2_v130 {
    function domainSeparator() external view returns (bytes32);
}

interface IMultiSendCallOnly_v130 {
    function multiSend(bytes memory transactions) external payable;
}

interface IMultiSend_v130 {
    function multiSend(bytes memory transactions) external payable;
}

interface IPermit2 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}

interface ISenderCreator_v060 {
    function createSender(bytes calldata initCode) external returns (address sender);
}

interface ISimpleAccountFactory {
    function getAddress(address owner, uint256 salt) external view returns (address);
}

interface IEntryPoint_v060 {
    function getSenderAddress(bytes calldata initCode) external;
}

interface ISenderCreator_v070 {
    function createSender(bytes calldata initCode) external returns (address sender);
}

interface IEntryPoint_v070 {
    function getSenderAddress(bytes calldata initCode) external;
}

interface IERC1820Registry {
    function interfaceHash(string calldata interfaceName) external pure returns (bytes32);
}
