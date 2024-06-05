import React from 'react'

import { getIcon } from '~/lib/sourceChains'
interface Props {
  chainId?: string
  name?: string
}

const RollupIcon: React.FC<Props> = ({ chainId, name }) => {
  return <img className='max-w-none rounded-full' src={getIcon(chainId || '') || ''} alt="" />
}

export default RollupIcon
