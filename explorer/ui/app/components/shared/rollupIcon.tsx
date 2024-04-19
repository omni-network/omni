import React from 'react'

import { getIcon } from '~/lib/sourceChains'
interface Props {
  chainId?: string
  name?: string
}

const RollupIcon: React.FC<Props> = ({ chainId, name }) => {
  return <img src={getIcon(chainId || '') || ''} alt="" />
}

export default RollupIcon
