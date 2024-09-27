# Walkthrough


In this guide we'll take a deeper look at the contracts that make up **[xstake](https://github.com/omni-network/xstake) - the simple chain-abstracted staking app** we deployed in the [Getting Started](/).


The app has two contracts:

- [`XStaker`](https://github.com/omni-network/xstake/blob/ad4cbb/src/XStaker.sol) - deployed on multiple chains, accepts ERC20 deposits.
- [`XStakeController`](https://github.com/omni-network/xstake/blob/ad4cbb/src/XStakeController.sol) - deployed on Omni, tracks stake across all chains.


### XApp

Both inherit from Omni's `XApp` utility base contract.

```solidity
contract XStaker is XApp {
  // ...
}
```

`XApp` makes it easy to send & receive cross-chain calls. Together, simple cross-chain calls + chain-aware global state compose the core of chain-abstracted applications built on Omni.


Using the xstake example, we'll walk through how to:

- Inform global state via cross-chain call (`xcall`)
- Read the context of an xcall when it's received (`xmsg`)


### xcall


Our [`XStaker`](https://github.com/omni-network/xstake/blob/ad4cbb/src/XStaker.sol) contracts needs to accept ERC20 deposits. Let's start with a simple function that does just that.

```solidity
function stake(uint256 amount) public payable {
    require(amount > 0, "XStaker: insufficient amount");
    require(token.transferFrom(msg.sender, address(this), amount), "XStaker: transfer failed");

    // now what?
}
```

If our app only lived on one chain, we'd track the user's deposit and be done. But our app spans multiple chains, with global state managed by the `XStakeController` on Omni. To update it, we'll use an `xcall`.


```solidity
// make a cross-chain call
xcall({
    to: controller,             // to the XStakeController
    destChainId: omniChainId(), // on Omni
    data: ?                     // calling which function?
    gasLimit: ?                 // with what gas limit?
});
```


### xmsg


We need a function to call on `XStakeController` that updates the global state. This function needs to be _aware_ of the cross-chain context in which it's called. This context is available in the `xmsg`.

```solidity
xmsg.sourceChainId  // where the xcall originated
xmsg.sender         // who sent it
```

To read this context, mark your receiving function with the `xrecv` modifier.


```solidity
function recordStake(address user, uint256 amount) public xrecv {
    // now read xmsg as needed
}
```

Using `xmsg`, we can authorize cross-chain calls to known `XStaker` deployments.


```solidity
// register XStaker deployments by chain id
mapping(uint64 => address) public xstakerOn;

function recordStake(address user, uint256 amount) external xrecv {
    require(xstakerOn[xmsg.sourceChainId] != address(0), "Controller: unsupported chain");
    require(xstakerOn[xmsg.sourceChainId] == xmsg.sender, "Controller: only xstaker");

    stakeOn[user][xmsg.sourceChainId] += amount;
}
```


### Putting it all together


With `XStakeController.recordStake` we have a function for our `XStaker` to call when a user makes a deposit.

```solidity
// XSaker.sol
function stake(uint256 amount) public payable {
    require(amount > 0, "XStaker: insufficient amount");
    require(token.transferFrom(msg.sender, address(this), amount), "XStaker: transfer failed");

    xcall({
        destChainId: omniChainId(),
        to: controller,
        data: abi.encodeCall(XStakeController.recordStake, (msg.sender, amount)),
        gasLimit: 100_000
    });
}

// XStakeController.sol
function recordStake(address user, uint256 amount) external xrecv {
    require(xstakerOn[xmsg.sourceChainId] != address(0), "Controller: unsupported chain");
    require(xstakerOn[xmsg.sourceChainId] == xmsg.sender, "Controller: only xstaker");

    stakeOn[user][xmsg.sourceChainId] += amount;
}
```

### Fees


Omni charge fees for each `xcall`. The `xcall(...)` utility will pay the fees if funds are avaible - either in the contract or sent with the call.
To make sure sufficient fees are sent with the call, add the following check:

```solidity
uint256 fee = xcall(...)
require(msg.value >= fee);
```

It's often useful to know the fee beforehand, so you can set the proper `msg.value`. For this, the `XStaker` contract has a `stakeFee` function, using the `feeFor` utility.

```solidity
function stakeFee(uint256 amount) public view returns (uint256) {
    return feeFor({
        destChainId: omniChainId(),
        data: abi.encodeCall(XStakeController.recordStake, (msg.sender, amount)),
        gasLimit: 100_000
    });
}
```

This is a useful pattern. For each `xcall` your contract makes, have a corresponding fee getter. This way, frontends, or other contracts, can set the correct `msg.value` with each call.


### Next Steps

This walkthrough uses code snippets without context, and omits some flows / concepts. Notably:

- unstaking
- xcall confirmation levels

The full code is available and generously commented [here](https://github.com/omni-network/xstake/tree/ad4cbb).
