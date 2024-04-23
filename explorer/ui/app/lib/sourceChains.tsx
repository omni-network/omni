import Arbiscan from '~/assets/images/Arbiscan.svg'
import Caldera from '~/assets/images/Caldera.svg'
import EigenLayer from '~/assets/images/EigenLayer.svg'
import Espresso from '~/assets/images/Espresso.svg'
import Linea from '~/assets/images/Linea.svg'
import Optimism from '~/assets/images/Optimism.svg'
import Polygon from '~/assets/images/Polygon.svg'
import Scroll from '~/assets/images/Scroll.svg'
import Generic from '~/assets/images/Generic.svg'

export const mappedSourceChains = (
  sourceChains: Array<{ ChainID: string; Name: string } | null>,
) => {
  if (sourceChains === null) {
    return []
  }

  return sourceChains.map(chain => ({
    ...chain,
    Icon: getIcon(chain?.ChainID || ''),
    DisplayName: chain?.Name.replaceAll('_', ' '),
    Urls: {
      txHash: getBaseUrl(chain?.ChainID || '', 'tx'),
      senderAddress: getBaseUrl(chain?.ChainID || '', 'senderAddress'),
      blockHash: getBaseUrl(chain?.ChainID || '', 'blockHash'),
      destHash: getBaseUrl(chain?.ChainID || '', 'destHash'),
    },
  }))
}

// this is hard coded on front-end until this data can come with the source chains from back-end
export const getBaseUrl = (chainId: string, type: 'tx' | 'senderAddress' | 'blockHash' | 'destHash') => {
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
