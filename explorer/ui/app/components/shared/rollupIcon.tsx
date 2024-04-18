import React from 'react'

import Arbiscan from '~/assets/images/Arbiscan.svg'
import Caldera from '~/assets/images/Caldera.svg'
import EigenLayer from '~/assets/images/EigenLayer.svg'
import Espresso from '~/assets/images/Espresso.svg'
import Linea from '~/assets/images/Linea.svg'
import Optimism from '~/assets/images/Optimism.svg'
import Polygon from '~/assets/images/Polygon.svg'
import Scroll from '~/assets/images/Scroll.svg'

interface Props {
  chainId?: string
  name?: string
}

const RollupIcon: React.FC<Props> = ({ chainId, name }) => {
  let image = Optimism
  switch (chainId) {
    case '0x89':
      image = Polygon
      break
    case '0xe708':
      image = Linea
      break
    case '0x82750':
      image = Scroll
      break
  }
  return <img src={image} alt="" />
}

export default RollupIcon
