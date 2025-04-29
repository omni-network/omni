const { ethers } = require('ethers')

// contract addresses - replace these with actual addresses
// use the proxy address for GenesisStake since it's upgradeable
const OMNI_TOKEN_ADDRESS = '0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07'
const GENESIS_STAKE_ADDRESS = '0x62335BbA8B27606B3CB70c818C996900479D1b8B'

// amount to stake in OMNI tokens
const STAKE_AMOUNT = '1' // 1 OMNI token

async function main() {
  // connect to the network
  const provider = new ethers.JsonRpcProvider('https://holesky.rpc.thirdweb.com')
  const wallet = new ethers.Wallet(
    'f4932749270a9750b621965b8b2213ae7a649d70301290c03e4cbd785338e695',
    provider,
  )

  // create contract instances
  const tokenContract = new ethers.Contract(
    OMNI_TOKEN_ADDRESS,
    [
      'function approve(address spender, uint256 amount) returns (bool)',
      'function allowance(address owner, address spender) view returns (uint256)',
    ],
    wallet,
  )

  const stakeContract = new ethers.Contract(
    GENESIS_STAKE_ADDRESS,
    ['function stake(uint256 amount)', 'function stakeFor(address recipient, uint256 amount)'],
    wallet,
  )

  // convert stake amount to wei
  const amount = ethers.parseEther(STAKE_AMOUNT)

  try {
    // check current allowance
    const currentAllowance = await tokenContract.allowance(wallet.address, GENESIS_STAKE_ADDRESS)
    console.log(`Current allowance: ${ethers.formatEther(currentAllowance)} OMNI`)

    // if allowance is less than amount, approve more
    if (currentAllowance < amount) {
      console.log('Approving tokens...')
      const approveTx = await tokenContract.approve(GENESIS_STAKE_ADDRESS, amount)
      await approveTx.wait()
      console.log('Approval successful!')
    }

    // stake tokens
    console.log('Staking tokens...')
    const stakeTx = await stakeContract.stake(amount)
    await stakeTx.wait()
    console.log('Staking successful!')
  } catch (error) {
    console.error('Error:', error.message)
  }
}

main().catch(error => {
  console.error(error)
  process.exit(1)
})
