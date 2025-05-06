import { useOrder, useQuote } from '@omni-network/react'
import { formatEther, parseEther, zeroAddress } from 'viem'
import { baseSepolia, holesky } from 'viem/chains'
import { useAccount, useConnect, useDisconnect, useSwitchChain } from 'wagmi'

function App() {
  return (
    <>
      <Account />
      <Order />
      <Quote />
    </>
  )
}

function Account() {
  const account = useAccount()
  const { connectors, connect } = useConnect()
  const { disconnect } = useDisconnect()

  return (
    <div>
      <h2>Account</h2>

      <div>
        account: {account.address}
        <br />
        chainId: {account.chainId}
        <br />
        status: {account.status}
      </div>

      {account.status !== 'disconnected' ? (
        <button type="button" onClick={() => disconnect()}>
          Disconnect
        </button>
      ) : (
        connectors.map((connector) => (
          <button
            key={connector.uid}
            onClick={() =>
              connect({
                connector,
              })
            }
            type="button"
          >
            {connector.name}
          </button>
        ))
      )}
    </div>
  )
}

function Quote() {
  const account = useAccount()
  const quote = useQuote({
    srcChainId: baseSepolia.id,
    destChainId: holesky.id,
    deposit: { amount: parseEther('0.1'), isNative: true },
    expense: { isNative: true },
    mode: 'expense',
    enabled: !!account?.address,
  })

  return (
    <div>
      <h2>Order</h2>
      {account?.address ? (
        <>
          <h4>Quote swap amount</h4>
          <div>isSuccess: {quote.isSuccess}</div>
          <div>isPending: {quote.isPending}</div>
          <div>isError: {quote.isError}</div>
          <div>
            quote.deposit.amount:{' '}
            {quote.isSuccess ? formatEther(quote?.deposit.amount) : ''}
          </div>
          <div>
            quote.expense.amount:{' '}
            {quote.isSuccess ? formatEther(quote?.expense.amount) : ''}
          </div>
        </>
      ) : (
        <div>connect...</div>
      )}
    </div>
  )
}

function Order() {
  const expectedSrcChainId = baseSepolia.id
  const account = useAccount()
  const { switchChain } = useSwitchChain()
  const order = useOrder({
    owner: account?.address,
    srcChainId: baseSepolia.id,
    destChainId: holesky.id,
    deposit: { amount: parseEther('0.1') },
    expense: {
      amount: parseEther('0.099'),
      spender: zeroAddress,
    },
    calls: [{ target: account.address ?? '0x', value: parseEther('0.099') }],
    validateEnabled: !!account?.address,
  })

  return (
    <div>
      <h2>Order</h2>
      {account?.address ? (
        <>
          <h4>Swap .1 eth from base sepolia to holesky</h4>
          <div>chainId: {account.chainId}</div>
          {expectedSrcChainId !== account.chainId && (
            <>
              <div>
                <div>wrong chain</div>
                <button
                  onClick={() => switchChain({ chainId: expectedSrcChainId })}
                  type="button"
                >
                  Switch chain
                </button>
              </div>
              <br />
            </>
          )}
          <div>validation: {order.validation?.status}</div>
          <div>status: {order.status}</div>
          <div>src chain tx hash: {order.txHash}</div>
          <div>isError: {order.isError}</div>
          <div>error: {order.error?.message}</div>
          <div>orderId: {order.orderId}</div>
          <button
            onClick={() => order.open()}
            disabled={
              order.validation?.status !== 'accepted' ||
              expectedSrcChainId !== account.chainId
            }
            type="button"
          >
            Swap
          </button>
        </>
      ) : (
        <div>connect...</div>
      )}
    </div>
  )
}

export default App
