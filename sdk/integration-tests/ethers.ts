import { ethers } from 'ethers'
import {
  type Account,
  type Address,
  type Chain,
  type CustomTransport,
  type EIP1193RequestFn,
  type PublicClient,
  RpcRequestError,
  type WalletClient,
  createPublicClient,
  createWalletClient,
  custom,
} from 'viem'

// Helper type guard to check if the object is an ethers v6 Signer
function isEthersSigner(
  providerOrSigner: ethers.Provider | ethers.Signer,
): providerOrSigner is ethers.Signer {
  return 'signTransaction' in providerOrSigner
}

// Define more precise return types based on input
type EthersToViemClientResult<
  TProviderOrSigner extends ethers.Provider | ethers.Signer,
> = TProviderOrSigner extends ethers.Signer
  ? WalletClient<CustomTransport, Chain | undefined, Account>
  : PublicClient<CustomTransport, Chain | undefined>

// Creates a viem Client instance from an ethers Provider or Signer
export async function ethersToViemClient<
  TProviderOrSigner extends ethers.Provider | ethers.Signer,
>(
  providerOrSigner: TProviderOrSigner,
  chain: Chain,
): Promise<EthersToViemClientResult<TProviderOrSigner>> {
  // Get the underlying provider
  const provider = isEthersSigner(providerOrSigner)
    ? providerOrSigner.provider
    : providerOrSigner

  if (!provider) {
    throw new Error(
      'Ethers Provider not found. Ensure the Signer is connected to a Provider.',
    )
  }

  // Define the Custom Transport using EIP-1193 request structure
  const request: EIP1193RequestFn = async ({ method, params }) => {
    // --- Handle WalletClient methods ---
    if (isEthersSigner(providerOrSigner)) {
      const signer = providerOrSigner
      try {
        if (method === 'eth_requestAccounts' || method === 'eth_accounts') {
          const address = await signer.getAddress()
          return [address as Address]
        }

        if (
          method === 'personal_sign' &&
          Array.isArray(params) &&
          params.length >= 2
        ) {
          const message = params[0] as string // Hex string message
          const address = params[1] as string // Address requested to sign
          const signerAddress = await signer.getAddress()
          if (address.toLowerCase() !== signerAddress.toLowerCase()) {
            throw new RpcRequestError({
              body: { method, params },
              error: {
                code: -32000,
                message: 'Requested address does not match signer',
              },
              url: '',
            })
          }
          // ethers v6 signMessage expects Uint8Array
          const messageBytes = ethers.getBytes(message)
          return await signer.signMessage(messageBytes)
        }

        if (
          method === 'eth_signTypedData_v4' &&
          Array.isArray(params) &&
          params.length >= 2
        ) {
          const address = params[0] as string // Address requested to sign
          const data = params[1] as string // JSON stringified typed data
          const signerAddress = await signer.getAddress()
          if (address.toLowerCase() !== signerAddress.toLowerCase()) {
            throw new RpcRequestError({
              body: { method, params },
              error: {
                code: -32000,
                message: 'Requested address does not match signer',
              },
              url: '',
            })
          }
          const { domain, types, message } = JSON.parse(data)
          // ethers v6 automatically adds the EIP712Domain type
          // delete EIP712Domain if it exists in types object
          if (types?.EIP712Domain) {
            types.EIP712Domain = undefined
          }
          return await signer.signTypedData(domain, types, message)
        }

        if (
          method === 'eth_sendTransaction' &&
          Array.isArray(params) &&
          params.length > 0
        ) {
          const tx = params[0] as ethers.TransactionRequest
          // Ensure gasLimit is used if provided as gas (common viem pattern)
          if ('gas' in tx && tx.gas !== undefined && !('gasLimit' in tx)) {
            tx.gasLimit = tx.gas as bigint | null | undefined
            ;(tx as { gas?: unknown }).gas = undefined
          }
          // ethers v6 requires 'to' to be null explicitly for contract creation
          if (tx.to === undefined) {
            tx.to = null
          }
          const submittedTx = await signer.sendTransaction(tx)
          return submittedTx.hash
        }
      } catch (error) {
        // Re-throw ethers errors as RpcRequestError for viem compatibility
        throw new RpcRequestError({
          body: { method, params },
          error: {
            code: (error as { code?: number })?.code ?? -32000,
            message:
              (error as { message?: string })?.message ?? 'Ethers Signer Error',
          },
          url: '', // No specific URL for local signer actions
        })
      }
    }

    // --- Handle PublicClient / Fallback methods ---
    // All other methods are forwarded to the underlying ethers provider
    try {
      const result: unknown = await provider.send(
        method,
        params as unknown[] | Record<string, unknown>,
      )
      return result
    } catch (error) {
      // Re-throw ethers provider errors as RpcRequestError
      throw new RpcRequestError({
        body: { method, params },
        error: {
          code: (error as { code?: number })?.code ?? -32603,
          message:
            (error as { message?: string })?.message ?? 'Ethers Provider Error',
        }, // -32603 Internal JSON-RPC error
        url: '', // RPC URL is not directly available from provider instance easily
      })
    }
  }

  const transport = custom({ request })

  // Create the appropriate viem Client
  if (isEthersSigner(providerOrSigner)) {
    // Create a WalletClient
    const signer = providerOrSigner
    const account: Account = {
      address: (await signer.getAddress()) as Address,
      type: 'json-rpc', // Indicates interaction via RPC requests handled by the signer
    }

    return createWalletClient({
      account,
      chain,
      transport,
    }) as EthersToViemClientResult<TProviderOrSigner>
  }

  // Create a PublicClient
  return createPublicClient({
    chain,
    transport,
  }) as EthersToViemClientResult<TProviderOrSigner>
}
