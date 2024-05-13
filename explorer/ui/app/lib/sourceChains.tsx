import Generic from '~/assets/images/Generic.svg'

export const mappedSourceChains = (
  sourceChains: Array<{ chainID: string; name: string; logoUrl: string; urls: object[] } | null>,
) => {
  if (sourceChains === null) {
    return []
  }

  return sourceChains.map(chain => {
    return {
      ...chain,
      Icon: chain?.logoUrl,
      DisplayName: chain?.name,
      urls: {
        txHash: getBaseUrl(chain?.chainID || '', 'tx'),
        senderAddress: getBaseUrl(chain?.chainID || '', 'senderAddress'),
        blockHash: getBaseUrl(chain?.chainID || '', 'blockHash'),
        destHash: getBaseUrl(chain?.chainID || '', 'destHash'),
      },
    }
  })
}

// this is hard coded on front-end until this data can come with the source chains from back-end
export const getBaseUrl = (
  chainId: string,
  type: 'tx' | 'senderAddress' | 'blockHash' | 'destHash',
) => {
  switch (chainId) {
    case '0x40b1': // mock
      switch (type) {
        case 'blockHash':
          return 'https://sepolia.arbiscan.io/block'
        case 'destHash':
          return 'https://sepolia.arbiscan.io/address'
        case 'senderAddress':
          return 'https://sepolia.arbiscan.io/address'
        case 'tx':
          return 'https://sepolia.arbiscan.io/tx'
        default:
          return '/'
      }
    case '0x64': // mock
      switch (type) {
        case 'blockHash':
          return 'https://sepolia.arbiscan.io/block'
        case 'destHash':
          return 'https://sepolia.arbiscan.io/address'
        case 'senderAddress':
          return 'https://sepolia.arbiscan.io/address'
        case 'tx':
          return 'https://sepolia.arbiscan.io/tx'
        default:
          return '/'
      }
    case '0xc8': // mock
      switch (type) {
        case 'blockHash':
          return 'https://sepolia.arbiscan.io/block'
        case 'destHash':
          return 'https://sepolia.arbiscan.io/address'
        case 'senderAddress':
          return 'https://sepolia.arbiscan.io/address'
        case 'tx':
          return 'https://sepolia.arbiscan.io/tx'
        default:
          return '/'
      }
    case '0x66eee': // Arbitrum Sepolia test net
      switch (type) {
        case 'blockHash':
          return 'https://sepolia.arbiscan.io/block'
        case 'destHash':
          return 'https://sepolia.arbiscan.io/address'
        case 'senderAddress':
          return 'https://sepolia.arbiscan.io/address'
        case 'tx':
          return 'https://sepolia.arbiscan.io/tx'
        default:
          return '/'
      }
    case '0xa5': // omni test net
      switch (type) {
        case 'blockHash':
          return 'https://omni-testnet.blockscout.com/block'
        case 'destHash':
          return 'https://sepolia.arbiscan.io/address'
        case 'senderAddress':
          return 'https://omni-testnet.blockscout.com/address'
        case 'tx':
          return 'https://omni-testnet.blockscout.com/tx'
        default:
          return '/'
      }
    case '0xaa37dc': // OP Sepolia Test net (Optimism)
      switch (type) {
        case 'blockHash':
          return 'https://sepolia-optimism.etherscan.io/block'
        case 'destHash':
          return 'https://sepolia.arbiscan.io/address'
        case 'senderAddress':
          return 'https://sepolia-optimism.etherscan.io/address'
        case 'tx':
          return 'https://sepolia-optimism.etherscan.io/tx'
        default:
          return '/'
      }
  }

  return '/'
}

export const getIcon = (ChainID: string) => {
  let icon = Generic
  switch (ChainID) {
    case '0x40b1':
      icon = Generic
      break
    case '0x64':
      icon = Generic
      break
    case '0xc8':
      icon = Generic
      break
    case '0x66eee': // Arbitrum Sepolia test net
      icon = Generic
      break
    case '0xa5': // omni test net
      icon = Generic
      break
  }
  return icon
}

export const getBlockUrl = (ChainID, blockHash, supportedChains) => {
  const chain = supportedChains.find(chain => chain.ChainID === ChainID)

  if (chain) {
    return `${chain.BaseExplorerUrl}/block/${blockHash}`
  }

  return '/'
}

export const getAddressUrl = (ChainID, addressHash, supportedChains) => {
  const chain = supportedChains.find(chain => chain.ChainID === ChainID)

  if (chain) {
    return `${chain.BaseExplorerUrl}/address/${addressHash}`
  }

  return '/'
}

export const getTxUrl = (ChainID, txHash, supportedChains) => {
  const chain = supportedChains.find(chain => chain.ChainID === ChainID)

  if (chain) {
    return `${chain.BaseExplorerUrl}/tx/${txHash}`
  }

  return '/'
}
