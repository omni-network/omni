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

    function delegateN(address validator, uint256 value, uint256 n) external payable {
        IStakingContract staking = IStakingContract(stakingContract);
        for (uint256 i = 0; i < n; i++) {
            staking.delegate{ value: value }(validator);
        }
    }

    function undelegateN(address validator, uint256 value, uint256 amount, uint256 n) external payable {
        IStakingContract staking = IStakingContract(stakingContract);
        for (uint256 i = 0; i < n; i++) {
            staking.undelegate{ value: value }(validator, amount);
        }
    }

    enum Method {
        Delegate,
        Undelegate
    }

    struct Call {
        Method method;
        uint256 value;
        address validator;
        uint256 amount;
    }

    function proxy(Call[] calldata calls) external payable {
        IStakingContract staking = IStakingContract(stakingContract);

        for (uint256 i = 0; i < calls.length; i++) {
            if (calls[i].method == Method.Delegate) {
                staking.delegate{ value: calls[i].value }(calls[i].validator);
            } else if (calls[i].method == Method.Undelegate) {
                staking.undelegate{ value: calls[i].value }(calls[i].validator, calls[i].amount);
            }
        }

        // Return any excess ETH
        uint256 remainingBalance = address(this).balance;
        if (remainingBalance > 0) {
            (bool success,) = msg.sender.call{ value: remainingBalance }("");
            require(success, "ETH refund failed");
        }
    }

    // Allow contract to receive ETH
    receive() external payable { }
}
