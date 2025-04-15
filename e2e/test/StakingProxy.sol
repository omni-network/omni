// This smart-contract batches calls to Staking.sol
pragma solidity ^0.8.0;

interface IStakingContract {
    function delegate(address validator) external payable;
    function undelegate(address validator, uint256 amount) external payable;
}

contract StakingProxy {
    address public stakingContract;

    constructor(address _stakingContract) {
        stakingContract = _stakingContract;
    }

    struct Call {
        string method;
        uint256 value;
        address validator;
        uint256 amount;
    }

    function proxy(Call[] calldata calls) external payable {
        uint256 totalValue = 0;
        for (uint i = 0; i < calls.length; i++) {
            totalValue += calls[i].value;
        }
        require(msg.value >= totalValue, "Insufficient ETH sent");

        IStakingContract staking = IStakingContract(stakingContract);

        for (uint i = 0; i < calls.length; i++) {
            if (keccak256(abi.encodePacked(calls[i].method)) == keccak256(abi.encodePacked("delegate"))) {
                staking.delegate{value: calls[i].value}(calls[i].validator);
            } else if (keccak256(abi.encodePacked(calls[i].method)) == keccak256(abi.encodePacked("undelegate"))) {
                staking.undelegate{value: calls[i].value}(calls[i].validator, calls[i].amount);
            }
        }

        // Return any excess ETH
        uint256 remainingBalance = address(this).balance;
        if (remainingBalance > 0) {
            (bool success, ) = msg.sender.call{value: remainingBalance}("");
            require(success, "ETH refund failed");
        }
    }

    // Allow contract to receive ETH
    receive() external payable {}
}
