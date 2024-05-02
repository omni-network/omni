# add support for ADD and OR predicates
type FilterParams implements Filter {
  "Source chain ID"
  sourceChainID: BigInt

  "Destination chain ID"
  destChainID: BigInt

  "Source or destination address"
  address: Address

  "Receipt transaction hash on the source or destination chain"
  txHash: Bytes32

  "Status of the message"
  status: XMsgStatus
}


union Filter = AndFilter | OrFilter

type AndFilter implements Filter
type OrFilter implements Filter



# union SearchResult = XMsgByAddr | XMsgByDest | XMsgBySrc | XMsgByID | XMsgByBlock | XMsgByTxHash
# XMsgByAddr - either from or to
# XMsgByTxHash - either source tx hash (the one that triggered xmsg) or receipt tx hash (receive tx hash)


# FE - use /?query=<search-term> - 
 - if no results - 404
 - if 1 result - redirect to it; otherwise
 - show a list of results with links to the actual pages

# partial search
/?query=0x12 - autocompletion (optional autocompletion)

# - endpoint for seach results (autocompletion or list of results)
# - endpoint for the actual redirect