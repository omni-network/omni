import { QueryClient } from '@tanstack/react-query'
import { queryOpts } from '../components/GitHubCodeBlock/useCodeBlock'

export const client = new QueryClient()

// prefectch known codeblock queries
const knownCodeBlocks = [
  'https://github.com/omni-network/omni/blob/main/contracts/src/interfaces/IOmniPortal.sol',
  'https://github.com/omni-network/omni/blob/main/contracts/src/pkg/XApp.sol',
  'https://github.com/omni-network/omni/blob/main/contracts/src/libraries/XTypes.sol',
  'https://github.com/omni-network/omni-forge-template/blob/main/src/XGreeter.sol',
  'https://github.com/omni-network/omni/blob/059303647e07fc3481e379b710922e2b84b1827f/contracts/src/pkg/XApp.sol#L56-L65',
  'https://github.com/omni-network/omni/blob/059303647e07fc3481e379b710922e2b84b1827f/contracts/src/protocol/OmniPortal.sol#L135-L151',
  'https://github.com/omni-network/omni/blob/059303647e07fc3481e379b710922e2b84b1827f/contracts/src/protocol/OmniPortal.sol#L165-L220',
]

const prefech = () =>
  knownCodeBlocks.forEach(url => client.prefetchQuery(queryOpts(url)))

prefech()
