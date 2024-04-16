import { graphql } from '~/graphql'

export const supportedchains = graphql(`
  query SupportedChains {
    supportedchains {
      ChainID
      Name
    }
  }
`)
