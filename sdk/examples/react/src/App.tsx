import { useOrder, useQuote } from '@omni-network/react'
import { formatEther, parseEther, zeroAddress } from 'viem'
import { baseSepolia, holesky } from 'viem/chains'
import { useAccount, useConnect, useDisconnect, useSwitchChain } from 'wagmi'

function App() {
  return (
    <>
      <Account />
      <Quote />
      <Order />
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
    deposit: { amount: parseEther('0.1') },
    mode: 'expense',
    enabled: true,
  })

  return (
    <div>
      <h2>Quote</h2>
      {account?.address ? (
        <>
          <h4>Quote swap amount</h4>
          <div>
            isSuccess: <code>{JSON.stringify(quote.isSuccess)}</code>
          </div>
          <div>
            isPending: <code>{JSON.stringify(quote.isPending)}</code>
          </div>
          <div>
            isError: <code>{JSON.stringify(quote.isError)}</code>
          </div>
          <div>
            quote.deposit.amount:{' '}
            {quote.isSuccess ? `${formatEther(quote.deposit.amount)} ETH` : ""}
          </div>
          <div>
            quote.expense.amount:{' '}
            {quote.isSuccess ? `${formatEther(quote.expense.amount)} ETH` : ""}
          </div>
        </>
      ) : (
        <div>Please connect wallet...</div>
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
      amount: parseEther('0.199'),
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
          <div>
            src chain tx hash:{''}
            {order.txHash && (
              <a
                target="_blank"
                rel="noopener noreferrer"
                href={`https://sepolia.basescan.org/tx/${encodeURIComponent(
                  order.txHash
                )}`}
              >
                <code>{order.txHash}</code>
              </a>
            )}
          </div>
          <div>isError: <code>{JSON.stringify(order.isError)}</code></div>
          <div>error: <pre>{order.error?.message}</pre></div>
          <div>orderId: <code>{order.orderId}</code></div>
          <div>destTxHash: {order.destTxHash && (
              <a
                target="_blank"
                rel="noopener noreferrer"
                href={`https://holesky.etherscan.io/tx/${encodeURIComponent(
                  order.destTxHash
                )}`}
              >
                <code>{order.destTxHash}</code>
              </a>
            )}</div>
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
        <div>Please connect wallet...</div>
      )}
    </div>
  )
}

export default App
