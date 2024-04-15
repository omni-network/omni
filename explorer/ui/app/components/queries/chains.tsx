import { useQuery } from 'urql'
import { graphql } from '~/graphql'
import { Chain } from '~/graphql/graphql'

export function GetBlocksInRange(from: number, to: number): Chain[] {
  const [result] = useQuery({
    query: supportedchains,
  })
  const { data, fetching, error } = result

  var rows: Chain[] = []

  data?.supportedchains.map((chain: any) => {
    let c = {
      ChainID: chain.ChainID,
      Name: chain.Name,
    }

    rows.push(c)
  })

  return rows
}

export const supportedchains = graphql(`
  query SupportedChains {
    supportedchains {
      ChainID
      Name
    }
  }
`)
