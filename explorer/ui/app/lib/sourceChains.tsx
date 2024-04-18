import Arbiscan from '~/assets/images/Arbiscan.svg'
import Caldera from '~/assets/images/Caldera.svg'
import EigenLayer from '~/assets/images/EigenLayer.svg'
import Espresso from '~/assets/images/Espresso.svg'
import Linea from '~/assets/images/Linea.svg'
import Optimism from '~/assets/images/Optimism.svg'
import Polygon from '~/assets/images/Polygon.svg'
import Scroll from '~/assets/images/Scroll.svg'

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
    BaseExplorerUrl: getBaseUrl(chain?.ChainID || ''),
  }))
}

const getBaseUrl = (chainId: string) => {
    return 'https://optimism-sepolia.blockscout.com'
}

export const getIcon = (ChainID: string) => {
  let icon = Optimism
  switch (ChainID) {
    case '0x40b1':
      icon = Polygon
      break
    case '0x64':
      icon = Linea
      break
    case '0xc8':
      icon = Scroll
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
